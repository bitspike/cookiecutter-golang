// Package cmd is responsible with building the structure of the
// command line interface (CLI) of this application.
//
// This package should contain all the globals of the package
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Manager struct {
	rootCmd *cobra.Command
}

// NewManager is the function which will need to  be called before the execution
// of the command line at the beginning of the application.
func NewManager() *Manager {
	cobra.OnInitialize()

	newManager := Manager{
		rootCmd: CreateRootCmd(),
	}
	newManager.rootCmd.AddCommand(CreateVersionCmd())
	// Add new commands here.

	return &newManager
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
// This function will exit the app up-on failure.
func (c *Manager) Execute() {
	if err := c.rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
