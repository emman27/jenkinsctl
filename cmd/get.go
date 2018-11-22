package cmd

import (
	"github.com/emman27/jenkinsutils/builds"
	"github.com/emman27/jenkinsutils/output"
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
	if len(args) == 0 {
		output.Print(builds.List())
	}
}
