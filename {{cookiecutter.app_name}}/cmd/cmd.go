// Package cmd is responsible with building the structure of the
// command line interface (CLI) of this application.
//
// This package should contain all the globals of the package
package cmd

import "github.com/spf13/cobra"

// Initialize is the function which will need to  be called before the execution
// of the command line at the beginning of the application.
func Initialize() {
	cobra.OnInitialize()
	rootCmd.AddCommand(versionCmd)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
