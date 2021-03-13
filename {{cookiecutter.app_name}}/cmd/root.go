package cmd

import (
	"github.com/spf13/cobra"
)

// cfgFile is the way viper configuration lib is working. Must be global.
var cfgFile string //nolint:unused,deadcode,varcheck,gochecknoglobals

// CreateRootCmd creates and returns the root cobra.Command object.
func CreateRootCmd() *cobra.Command {
	// nolint:exhaustivestruct
	return &cobra.Command{
		Use:   "generated code example",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
	examples and usage of using your application. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		//      Run: func(cmd *cobra.Command, args []string) { },
	}
}
