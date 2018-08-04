package cmd

import (
	"github.com/ironstar-io/tokaido/services/docker"
	"github.com/ironstar-io/tokaido/utils"

	"github.com/spf13/cobra"
)

// LogsCmd - `tok logs [x]`
var LogsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Output container logs to the console",
	Long:  "Output container logs to the console for all or a single container. Aliases `docker-compose -f x logs [container]`",
	Run: func(cmd *cobra.Command, args []string) {
		utils.CheckCmdHard("docker-compose")

		docker.PrintLogs(args)
	},
}
