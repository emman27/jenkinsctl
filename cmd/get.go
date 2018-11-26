package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/emman27/jenkinsctl/output"
	"github.com/emman27/jenkinsctl/pkg/api/builds"
	"github.com/emman27/jenkinsctl/pkg/api/parameters"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a resource from Jenkins",
}

var build = &cobra.Command{
	Use:   "builds",
	Short: "Get a build from Jenkins",
	Run:   buildsCmd,
}

var parametersCommand = &cobra.Command{
	Use:   "parameters",
	Short: "Get build parameters from Jenkins",
	Run:   parametersCmd,
}

func init() {
	getCmd.AddCommand(build)
	getCmd.AddCommand(parametersCommand)
}

func buildsCmd(cmd *cobra.Command, args []string) {
	if len(args) == 1 {
		jobName := args[0]
		buildData, err := builds.List(client, jobName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		output.Print(buildData)
	} else if len(args) == 2 {
		jobName := args[0]
		buildID, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		buildData, err := builds.Get(client, jobName, buildID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		output.Print(&builds.Builds{*buildData})
	}
}

func parametersCmd(cmd *cobra.Command, args []string) {
	var (
		buildID int
		err     error
		params  *parameters.Parameters
	)
	if len(args) != 2 {
		fmt.Println("This command requires 2 parameters. Usage: jenkinsctl get parameters <jobname> <buildID>")
		os.Exit(1)
	}
	jobName := args[0]
	if buildID, err = strconv.Atoi(args[1]); err != nil {
		fmt.Printf("Build ID %v is not an integer]\n", args[1])
		os.Exit(1)
	}
	if params, err = parameters.Get(client, jobName, buildID); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	output.Print(params)
}
