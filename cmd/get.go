package cmd

import (
	"fmt"
	"os"

	"github.com/emman27/jenkinsutils/output"
	"github.com/emman27/jenkinsutils/pkg/api/builds"
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
	if len(args) != 0 {
		jobName := args[0]
		buildData, err := builds.List(client, jobName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		output.Print(buildData)
	}
}
