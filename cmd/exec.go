package cmd

import (
	"github.com/ironstar-io/tokaido/system/ssh"
	"github.com/ironstar-io/tokaido/utils"

	"fmt"

	"github.com/spf13/cobra"
)

// ExecCmd - `tok exec`
var ExecCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute a command inside your Drush container",
	Long:  "Execute a command inside your Drush container using SSH. Aliases `ssh <project-name>.tok -C command`",
	Run: func(cmd *cobra.Command, args []string) {
		utils.CheckCmdHard("docker-compose")

		r := ssh.ConnectCommand(args)
		fmt.Println(r)
	},
}
