package tasks

import (
	"fmt"
	"log"

	"github.com/plexcorp-pty-ld/tuxdeploy/core"
	"github.com/plexcorp-pty-ld/tuxdeploy/scripts"
	"github.com/plexcorp-pty-ld/tuxdeploy/termio"
)

func getSteps(config *core.AppConfig) []core.BuildStep {
	var steps []core.BuildStep
	steps = append(steps, core.BuildStep{
		StepName: "Install APT Packages", Code: scripts.INIT_SERVER, RunLocal: false},
	)

	steps = append(steps, core.BuildStep{
		StepName: "Setup SSH User", Code: scripts.GetSshUserSetup(config), RunLocal: false},
	)

	steps = append(steps, core.BuildStep{
		StepName: "Security Post Steps", Code: scripts.GetSecurityPostSteps(config), RunLocal: false},
	)

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

	for stepNum, step := range steps {

		response, err := core.RunSshCommand(client, &step, stepNum)

		if err != nil {
			termio.PrintError("Oops, step: "+step.StepName+" has errored with message: "+err.Error()+".", 60)
			return
		}

		stepNum++
		fmt.Println()
		termio.PrintRegularText(">>> BUILD LOG - "+step.StepName+" <<<<", 0)
		fmt.Println(">>>>>>>>>")
		fmt.Println("Ran: " + step.Code)
		fmt.Println(">>>>>>>>>")

		fmt.Println(string(response))

		termio.PrintRegularText(">>> END BUILD LOG - "+step.StepName+" <<<<", 0)
		fmt.Println()

	}

	msg := termio.GetColorText(termio.SUCCESS_TEXT, termio.SUCCESS_TEXT)
	fmt.Println(msg.Render("Yay! all steps for the intial server setup are complete."))
}
