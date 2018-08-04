package daemon

import (
	"github.com/ironstar-io/tokaido/conf"
	"github.com/ironstar-io/tokaido/system/console"
	"github.com/ironstar-io/tokaido/utils"

	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
)

type service struct {
	ProjectName string
	ProjectPath string
	Username    string
}

// LoadService ...
func LoadService(servicePath string) {
	_, err := utils.CommandSubSplitOutput("launchctl", "load", servicePath)
	if err != nil {
		log.Fatal("Unable to load sync service: ", err)
	}
}

// UnloadService ...
func UnloadService(servicePath string) {
	_, err := utils.CommandSubSplitOutput("launchctl", "unload", servicePath)
	if err != nil {
		log.Fatal("Unable to unload sync service: ", err)
	}
}

// StartService ...
func StartService(serviceName string) {
	_, err := utils.CommandSubSplitOutput("launchctl", "start", serviceName)
	if err != nil {
		log.Fatal("Unable to start the sync service: ", err)
	}
}

// StopService ...
func StopService(serviceName string) {
	_, err := utils.CommandSubSplitOutput("launchctl", "stop", serviceName)
	if err != nil {
		log.Fatal("Unable to stop the sync service: ", err)
	}
}

// DeleteService ...
func DeleteService(servicePath string) {
	os.Remove(servicePath)
}

// ServiceStatus checks if the unison background process is running
func ServiceStatus(serviceName string) string {
	c := conf.GetConfig()

	u, uErr := user.Current()
	if uErr != nil {
		log.Fatal(uErr)
	}

	o, _ := utils.CommandSubSplitOutput("launchctl", "print", "gui/"+u.Uid+"/"+serviceName)

	if c.Tokaido.Debug == true {
		fmt.Printf("\033[33m%s\033[0m\n", o)
	}

	if strings.Contains(o, "state = running") == true {
		return "running"
	}

	return "stopped"
}

// KillService ...
func KillService(serviceName string, servicePath string) {
	ps, _ := utils.CommandSubSplitOutput("launchctl", "list", serviceName)
	if ps != "" {
		console.Println(`
🔄  Removing the background sync process
	`, "")
		StopService(serviceName)
		UnloadService(servicePath)
		DeleteService(servicePath)
	}
}
