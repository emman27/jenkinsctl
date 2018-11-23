package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/emman27/jenkinsctl/output"
	"github.com/emman27/jenkinsctl/pkg/api/builds"
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

func init() {
	getCmd.AddCommand(build)
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
