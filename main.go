package main

import (
	"fmt"
	"os"

	"github.com/plexcorp-pty-ld/tuxdeploy/core"
	"github.com/plexcorp-pty-ld/tuxdeploy/tasks"
	"github.com/plexcorp-pty-ld/tuxdeploy/termio"
)

func main() {

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get current working directory: ", err)
		return
	}

	_, err = os.Stat(cwd + "/.tuxdeploy")
	if err != nil {

		os.Mkdir(cwd+"/.tuxdeploy", 0755)
	}

	appConfig := core.GetTomlConfig()

	if appConfig.Server.Address == "" {
		termio.DrawTerminalHeader("Application TOML file missing.")
		termio.PrintAppConfigNotFound()
		return
	}

	if appConfig.Server.NewSSHPort == 0 || appConfig.Server.NewSSHPort == appConfig.Server.Port {
		termio.DrawTerminalHeader("Application TOML file is invalid.")
		termio.PrintError("New SSH port is invalid or cannot be the same as your current SSH port.", 0)
		return
	}

	if len(appConfig.Server.NewUsername) < 4 || appConfig.Server.NewUsername == appConfig.Server.Username {
		termio.DrawTerminalHeader("Application TOML file is invalid.")
		termio.PrintError("New SSH username cannot be the same as your current username and needs to be at least 4 characters long.", 0)
		return
	}

	if appConfig.Server.NewUsername == "root" || appConfig.Server.NewUsername == "Root" {
		termio.DrawTerminalHeader("Application TOML file is invalid.")
		termio.PrintError("Root cannot be used as the new username.", 0)
		return
	}

	if len(os.Args) == 1 {
		termio.DrawMainMenu()
		return
	}

	switch os.Args[1] {

	case "build":
		tasks.BuildServer(appConfig)
	case "django":
		tasks.BuildServerDjango(appConfig)
	default:
		termio.DrawMainMenu()
	}

}
