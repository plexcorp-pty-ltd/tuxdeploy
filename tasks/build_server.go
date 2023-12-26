package tasks

import (
	"fmt"
	"log"
	"os"

	"github.com/adhocore/chin"
	"github.com/plexcorp-pty-ld/tuxdeploy/core"
	"github.com/plexcorp-pty-ld/tuxdeploy/scripts"
	"github.com/plexcorp-pty-ld/tuxdeploy/termio"
)

func getSteps(config *core.AppConfig) map[string]string {
	steps := make(map[string]string)
	steps["Install APT Packages"] = scripts.INIT_SERVER
	steps["Setup SSH User"] = scripts.GetSshUserSetup(config)
	steps["Security Post Steps"] = scripts.GetSecurityPostSteps(config)

	return steps
}

func BuildServer(config *core.AppConfig) {

	termio.DrawTerminalHeader("Ubuntu standard server setup. This utility will provision " +
		"a new server which includes SSH hardening and setting up a firewall.")

	sshConfig := core.GetIntialSSHConfig(config)
	client, err := core.GetSshClient(config, sshConfig)

	if err == nil {
		defer client.Close()
	} else {
		log.Fatal("Sorry, there was a problem connecting to your server:", err)
		return
	}

	steps := getSteps(config)
	tmpPath := GetTmpPath("initial_builder")
	tmpPathExists, lastStep := GetLastStepFromTmpFile(tmpPath)

	stepNum := 1

	for name, code := range steps {
		if lastStep >= stepNum {
			fmt.Printf("Skipping already done step %s\n", name)
			continue
		}

		spinner := chin.New()
		go spinner.Start()
		termio.PrintRegularText("Running step: "+name, 0)
		response, err := client.Run(code)
		spinner.Stop()

		if err != nil {
			termio.PrintError("Oops, step: "+name+" has errored with message: "+err.Error()+".", 60)
			return
		}

		stepNum++
		WriteStepToTmpFile(tmpPath, tmpPathExists, stepNum)
		fmt.Println()
		termio.PrintRegularText(">>> BUILD LOG - "+name+" <<<<", 0)
		fmt.Println(">>>>>>>>>")
		fmt.Println("Ran: " + code)
		fmt.Println(">>>>>>>>>")

		fmt.Println(string(response))

		termio.PrintRegularText(">>> END BUILD LOG - "+name+" <<<<", 0)
		fmt.Println()

	}

	os.Remove(tmpPath)

	msg := termio.GetColorText("#0B832C")
	fmt.Println(msg.Render("Yay! all steps for the intial server setup are complete."))
}
