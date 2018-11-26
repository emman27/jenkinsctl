package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version of the jenkinsctl command
const Version = "1.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version number for Jenkins CLI",
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}

func printVersion() {
	fmt.Println(Version)
}
