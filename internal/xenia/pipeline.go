package xenia

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/ardanlabs/kit/log"
	"github.com/coralproject/shelf/internal/platform/db"
	"github.com/coralproject/shelf/internal/platform/db/mongo"
	"github.com/coralproject/shelf/internal/wire"
	"github.com/coralproject/shelf/internal/xenia/query"
	"github.com/pborman/uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// execPipeline executes the sepcified pipeline query.
func execPipeline(context interface{}, db *db.DB, q *query.Query, vars map[string]string, data map[string]interface{}, explain bool) (docs, []map[string]interface{}, error) {

	// I am returning commands as the second return value because if there
	// is an error I need to send how far we got back to the client. If not,
	// the user will not understand the error message.

	// We need to check to see if the last command is the extended $save command.
	commands := q.Commands
	l := len(q.Commands) - 1

	// Validate we have scripts to run.
	if l < 0 {
		return docs{}, commands, errors.New("Invalid pipeline script")
	}

	// If the last command is a $save, capture its value and remove
	// it from the pipeline.
	var save map[string]interface{}
	if v, exists := q.Commands[l]["$save"]; exists {
		if cmd, ok := v.(map[string]interface{}); ok {
			save = cmd
		}

		commands = q.Commands[0:l]
	}

	var agg string
	var pipeline []bson.M

	// Iterate over the commands and build the pipeline.
	for _, command := range commands {

		// Do we have variables to be substitued.
		if vars != nil {
			if err := ProcessVariables(context, command, vars, data); err != nil {
				return docs{}, commands, err
			}
		}

		// Add the operation to the slice for the pipeline.
		pipeline = append(pipeline, command)

		// Build a logable version of this pipeline.
		agg += mongo.Query(command) + ",\n"
	}

	// Are we being asked to execute the query on a view.
	if q.Collection == "view" {

		// Materialize the view for the query.
		viewCol, err := materializeView(context, db, q, vars)
		if err != nil {
			return docs{}, commands, err
		}

		// Defer clean up of the temporary collection.
		defer cleanupView(context, db, viewCol)
	}

	// Do we want the explain output.
	if explain {

		// Build the pipeline function for the execution for explain.
		var m bson.M
		f := func(c *mgo.Collection) error {
			log.Dev(context, "executePipeline", "MGO Explain :\ndb.%s.aggregate([\n%s])", c.Name, agg)
			return c.Pipe(pipeline).Explain(&m)
		}

		// Execute the pipeline.
		if err := db.ExecuteMGO(context, q.Collection, f); err != nil {
			return docs{}, commands, err
		}

		return docs{q.Name, []bson.M{m}}, commands, nil
	}

	// Set the default timeout for the session.
	timeout := 25 * time.Second
	if q.Timeout != "" {
		if d, err := time.ParseDuration(q.Timeout); err != nil {
			log.Dev(context, "executePipeline", "WARNING : Unable to Set Timeout[%s], using default.", q.Timeout)
		} else {
			timeout = d
		}
	}

	log.Dev(context, "executePipeline", "MGO Timeout Set[%s]", timeout)

	// Build the pipeline function for the execution.
	var results []bson.M
	f := func(c *mgo.Collection) error {
		log.Dev(context, "executePipeline", "MGO Started\ndb.%s.aggregate([\n%s])", c.Name, agg)
		err := c.Pipe(pipeline).All(&results)
		return err
	}

	// Set the channel to one because we might not be around
	// waiting for the result on timeouts.
	wait := make(chan error, 1)

	// Execute the pipeline.
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Dev(context, "executePipeline", "******> Recovered from timing out")
			}
			log.Dev(context, "executePipeline", "MGO Response Complete")
		}()

		wait <- db.ExecuteMGOTimeout(context, timeout, q.Collection, f)
	}()

	// Did any errors occur.
	select {

	// Wait for the response from executing the pipeline.
	case err := <-wait:
		if err != nil {
			if _, ok := err.(*net.OpError); ok {
				log.Error(context, "executePipeline", err, "Timed out Network")
				return docs{}, commands, errors.New("Completed : Timed out executing commands")
			}

			log.Error(context, "executePipeline", err, "Completed")
			return docs{}, commands, err
		}

	// Wait to timeout the entire operation.
	case <-time.After(timeout):
		err := errors.New("Timedout executing commands")
		log.Error(context, "executePipeline", err, "Completed : Timed out Processing")
		return docs{}, commands, err
	}

	log.Dev(context, "executePipeline", "Completed")

	// If there were no results, return an empty array.
	if results == nil {
		return docs{q.Name, []bson.M{}}, commands, nil
	}

	// Perform any masking that is required.
	if err := processMasks(context, db, q.Collection, results); err != nil {
		return docs{}, commands, err
	}

	// Do we need to save the result.
	if save != nil {
		if err := saveResult(context, save, results, data); err != nil {
			return docs{}, commands, err
		}
	}

	return docs{q.Name, results}, commands, nil
}

// materializeView executes a view, creates a temporary collection for the view, and
// modifies the query to query the temporary collection.
func materializeView(context interface{}, db *db.DB, q *query.Query, vars map[string]string) (string, error) {

	// Make sure we have a valid connection to the graph.
	graph, err := db.GraphHandle(context)
	if err != nil {
		return "", err
	}

	// Make sure we have the information we need to execute the view.
	viewName, ok := vars["view"]
	if !ok {
		return "", fmt.Errorf("Vars does not include \"view\".")
	}

	itemKey, ok := vars["item"]
	if !ok {
		return "", fmt.Errorf("Vars does not include \"item\".")
	}

	// Generate a unique name for the collection.
	viewCol := uuid.New()

	// Prepare the parameters for executing the view.
	viewParams := wire.ViewParams{
		ViewName:          viewName,
		ItemKey:           itemKey,
		ResultsCollection: viewCol,
	}

	// Execute the view.
	if _, err := wire.Execute(context, db, graph, &viewParams); err != nil {
		return viewCol, err
	}

	// Provide the query with the temporary collection name.
	q.Collection = viewCol

	return viewCol, nil
}

// cleanupView removed a tempory view collection.
func cleanupView(context interface{}, db *db.DB, viewCol string) error {
	c, err := db.CollectionMGO(context, viewCol)
	if err != nil {
		return err
	}

	if err := c.DropCollection(); err != nil {
		return err
	}

	return nil
}

// saveResult processes the $save command for this result.
func saveResult(context interface{}, save map[string]interface{}, results []bson.M, data map[string]interface{}) error {

	// {"$map": "list"}

	// Capture the key and value and process the save.
	for cmd, value := range save {
		name, ok := value.(string)
		if !ok {
			err := fmt.Errorf("Save key \"%v\" is a %T but must be a string", value, value)
			log.Error(context, "saveResult", err, "Extracting save key")
			return err
		}

		switch cmd {

		// Save the results into the map under the specified key.
		case "$map":
			log.Dev(context, "saveResult", "Saving result to map[%s]", name)
			data[name] = results
			return nil

		default:
			err := fmt.Errorf("Invalid save location %q", cmd)
			log.Error(context, "saveResult", err, "Nothing saved")
			return err
		}
	}

	err := errors.New("Missing save document")
	log.Error(context, "saveResult", err, "Nothing saved")
	return err
}
