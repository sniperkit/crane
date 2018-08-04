package cmd

import (
	"github.com/ironstar-io/tokaido/services/docker"
	"github.com/ironstar-io/tokaido/utils"

	"github.com/spf13/cobra"
)

// PsCmd - `tok ps`
var PsCmd = &cobra.Command{
	Use:   "ps",
	Short: "docker ps for your containers",
	Long:  "Alias for running `docker-compose -f <composefile> ps`",
	Run: func(cmd *cobra.Command, args []string) {
		utils.CheckCmdHard("docker-compose")

		docker.HardCheckTokCompose()

		docker.Ps()
	},
}
