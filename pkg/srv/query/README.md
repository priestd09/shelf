
# query
    import "github.com\coralproject\shelf\pkg\srv\query"

Package query provides API's for managing querysets which will be used in
executing different aggregation tests against their respective data collection.

QuerySet
In query, records are required to follow specific formatting and are at this
point, only allowed to be in a json serializable format which meet the query.Set
structure.

The query set execution supports the following types:


        - Pipeline
          Pipeline query set types take advantage of MongoDB's aggregation API
        (the currently supported data backend), which allows insightful use of its
        internal query language, in providing context against data sets within the database.

        - Template

QuerySet Sample:
{


        "name":"spending_advice",
        "description":"tests against user spending rate and provides adequate advice on saving more",
        "enabled": true,
        "params":[
          {
            "name":"user_id",
            "default":"396bc782-6ac6-4183-a671-6e75ca5989a5",
            "desc":"provides the user_id to check against the collection"
          }
        ],
        "rules": [
        {
          "desc":"match spending rate over 20 dollars",
          "type":"pipeline",
          "continue": true,
          "script_options": {
            "collection":"demo_user_transactions",
            "has_date":false,
            "has_objectid": false
          },
          "save_options": {
            "save_as":"high_dollar_users",
            "variables": true,
            "to_json": true
          },
          "var_options":{},
          "scripts":[
            "{ \"$match\" : { \"user_id\" : \"#userId#\", \"category\" : \"gas\" }}",
            "{ \"$group\" : { \"_id\" : { \"category\" : \"$category\" }, \"amount\" : { \"$sum\" : \"$amount\" }}}",
            "{ \"$match\" : { \"amount\" : { \"$gt\" : 20.00} }}"
          ]
         }]

}






## func CreateSet
``` go
func CreateSet(context interface{}, ses *mgo.Session, qs Set) error
```
CreateSet is used to create Set documents in the db.


## func DeleteSet
``` go
func DeleteSet(context interface{}, ses *mgo.Session, name string) error
```
DeleteSet is used to remove an existing Set document.


## func GetSetNames
``` go
func GetSetNames(context interface{}, ses *mgo.Session) ([]string, error)
```
GetSetNames retrieves a list of rule names.


## func UpdateSet
``` go
func UpdateSet(context interface{}, ses *mgo.Session, qs Set) error
```
UpdateSet is used to update an existing Set document.



## type Query
``` go
type Query struct {
    Description   string        `bson:"desc,omitempty" json:"desc,omitempty"`                     // Description of this specific query.
    Type          string        `bson:"type" json:"type"`                                         // variable, inventory, pipeline, template
    Continue      bool          `bson:"continue,omitempty" json:"continue,omitempty"`             // Indicates that on failure to process the next query.
    ScriptOptions *ScriptOption `bson:"script_options,omitempty" json:"script_options,omitempty"` // Options associated with script processing.
    SaveOptions   *SaveOption   `bson:"save_options,omitempty" json:"save_options,omitempty"`     // Options associated with saving the result.
    VarOptions    *VarOption    `bson:"var_options,omitempty" json:"var_options,omitempty"`       // Options associated with variable processing.
    Scripts       []string      `bson:"scripts" json:"scripts"`                                   // Scripts to process for the query.
}
```
Query contains the configuration details for a query.
Options use a pointer so they can be excluded when not in use.











## type Result
``` go
type Result struct {
    FeedName   string      `json:"feed_name"`
    Collection string      `json:"collection"`
    QueryType  string      `json:"query_type"`
    Results    interface{} `json:"results"`
    Valid      bool        `json:"valid"`
    Error      bool        `json:"-"`
}
```
Result contains the result of an query set execution.











## type SaveOption
``` go
type SaveOption struct {
    SaveAs    string `bson:"save_as,omitempty" json:"save_as,omitempty"`     // Name of the memory variable to store the result into.
    Variables bool   `bson:"variables,omitempty" json:"variables,omitempty"` // Indicates if the result should be saved into the variables.
    ToJSON    bool   `bson:"to_json,omitempty" json:"to_json,omitempty"`     // Convert the string result to JSON. Template oriented.
}
```
SaveOption contains options for saving results.











## type ScriptOption
``` go
type ScriptOption struct {
    Collection  string `bson:"collection,omitempty" json:"collection,omitempty"`     // Name of the collection to use for processing the query.
    HasDate     bool   `bson:"has_date,omitempty" json:"has_date,omitempty"`         // Indicates there is a date to be pre-processed in the scripts.
    HasObjectID bool   `bson:"has_objectid,omitempty" json:"has_objectid,omitempty"` // Indicates there is an ObjectId to be pre-processed in the scripts.
}
```
ScriptOption contains options for processing the scripts.











## type Set
``` go
type Set struct {
    Name        string     `bson:"name" json:"name"`       // Name of the query set.
    Description string     `bson:"desc" json:"desc"`       // Description of the query set.
    Enabled     bool       `bson:"enabled" json:"enabled"` // If the query set is enabled to run.
    Params      []SetParam `bson:"params" json:"params"`   // Collection of parameters.
    Queries     []Query    `bson:"queries" json:"queries"` // Collection of queries.
}
```
Set contains the configuration details for a rule set.









### func GetSetByName
``` go
func GetSetByName(context interface{}, ses *mgo.Session, name string) (*Set, error)
```
GetSetByName retrieves the configuration for the specified Set.




## type SetParam
``` go
type SetParam struct {
    Name    string `bson:"name" json:"name"`       // Name of the parameter.
    Default string `bson:"default" json:"default"` // Default value for the parameter.
    Desc    string `bson:"desc" json:"desc"`       // Description about the parameter.
}
```
SetParam contains meta-data about a required parameter for the query.











## type VarOption
``` go
type VarOption struct {
    ObjectID bool `bson:"object_id,omitempty" json:"object_id,omitempty"` // Indicates to save ObjectId values with ObjectId tag.
}
```
VarOption contains options for processing variables.




- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)