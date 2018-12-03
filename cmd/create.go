package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Jenkins resource",
}

var createBuild = &cobra.Command{
	Use:   "build",
	Short: "Create a Jenkins build",
	Run:   createBuildCmd,
}

func init() {
	createCmd.AddCommand(createBuild)
}

func createBuildCmd(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("This command only takes one parameter. Usage: jenkinsctl create build hello_world")
		os.Exit(1)
	}
	fmt.Println("Build created")
}
