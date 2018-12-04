package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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

var params = pflag.StringSliceP("param", "p", []string{}, "jenkinsctl create build my_job --param something=value --param key=othervalue")

func init() {
	createCmd.AddCommand(createBuild)
}

func createBuildCmd(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("This command only takes one parameter. Usage: jenkinsctl create build hello_world")
		os.Exit(1)
	}
	jobName := args[0]
	parsed := map[string]string{}
	for _, param := range *params {
		arr := strings.Split(param, "=")
		if len(arr) != 2 {
			fmt.Printf("Parameter %v is invalid, should be key=value\n", param)
			os.Exit(1)
		}
		parsed[arr[0]] = arr[1]
	}
	_, err := client.CreateBuild(jobName, parsed)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
