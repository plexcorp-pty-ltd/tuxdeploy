package main

import (
	"os"

	"github.com/plexcorp-pty-ld/tuxdeploy/core"
	"github.com/plexcorp-pty-ld/tuxdeploy/tasks"
	"github.com/plexcorp-pty-ld/tuxdeploy/termio"
)

func main() {

	appConfig := core.GetTomlConfig()

	if appConfig.Server == "" {
		termio.DrawTerminalHeader("Application TOML file missing.")
		termio.PrintAppConfigNotFound(core.GetGlobalAppConfigPath())
		return
	}

	if appConfig.NewSSHPort == 0 || appConfig.NewSSHPort == appConfig.Port {
		termio.DrawTerminalHeader("Application TOML file is invalid.")
		termio.PrintError("New SSH port is invalid or cannot be the same as your current SSH port.", 0)
		return
	}

	if len(appConfig.NewUsername) < 4 || appConfig.NewUsername == appConfig.Username {
		termio.DrawTerminalHeader("Application TOML file is invalid.")
		termio.PrintError("New SSH username cannot be the same as your current username and needs to be at least 4 characters long.", 0)
		return
	}

	if appConfig.NewUsername == "root" || appConfig.NewUsername == "Root" {
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
	default:
		termio.DrawMainMenu()
	}

}
