package cmd

import (
	"github.com/ironstar-io/tokaido/services/docker"
	"github.com/ironstar-io/tokaido/services/unison"
	"github.com/ironstar-io/tokaido/system/console"
	"github.com/ironstar-io/tokaido/utils"

	"fmt"

	"github.com/spf13/cobra"
)

// DestroyCmd - `tok destroy`
var DestroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Stop and destroy all containers",
	Long:  "Gracefully stop and destroy your Tokaido containers - `docker-compose down`",
	Run: func(cmd *cobra.Command, args []string) {
		utils.CheckCmdHard("docker-compose")

		docker.HardCheckTokCompose()

		confirmDestroy := utils.ConfirmationPrompt(`🔥  This will also destroy the database inside your Tokaido environment. Are you sure?`, "n")
		if confirmDestroy == false {
			console.Println(`🍵  Exiting without change`, "")
			return
		}
		fmt.Println()

		docker.Down()

		unison.StopSyncService()
	},
}
