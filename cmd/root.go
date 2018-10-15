package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "protoapi",
	Short: "protoapi is a code generation command line tool",
	Long: `protoapi is a tool/document to help generate API code:
			1. as a IDL, to document API request/response and function for developers' reference;
			2. generate boilderplate code to save time and cost for developers`,
	Version: "0.1.1",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.SetUsageFunc(usageFunc)
	RootCmd.SetHelpTemplate(`{{.UsageString}}`)
}
