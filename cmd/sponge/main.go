// This program provides a set of commands for item functionality.
package main

import (
	"os"

	"github.com/ardanlabs/kit/cfg"
	"github.com/ardanlabs/kit/log"
	"github.com/coralproject/shelf/cmd/sponge/cmditem"
	"github.com/coralproject/shelf/cmd/sponge/cmdpattern"
	"github.com/spf13/cobra"
)

const (

	// cfgLoggingLevel is the key for the logging level.
	cfgLoggingLevel = "LOGGING_LEVEL"

	// cfgWebHost is the key for the web host.
	cfgWebHost = "WEB_HOST"
)

// wire includes information about the sponge cobra command.
var sponge = &cobra.Command{
	Use:   "sponge",
	Short: "Sponge provides the central cli housing of various cli tools that interface with the internal sponge API",
}

func main() {

	// Initialize the configuration
	if err := cfg.Init(cfg.EnvProvider{Namespace: "SPONGE"}); err != nil {
		sponge.Println("Unable to initialize configuration")
		os.Exit(1)
	}

	// Initialize the logging
	logLevel := func() int {
		ll, err := cfg.Int(cfgLoggingLevel)
		if err != nil {
			return log.NONE
		}
		return ll
	}

	log.Init(os.Stderr, logLevel, log.Ldefault)
	sponge.Println("Using log level", logLevel())

	// Add the item commands to the CLI tool.
	sponge.AddCommand(
		cmditem.GetCommands(),
		cmdpattern.GetCommands(),
	)

	// Execute the command.
	sponge.Execute()
}
