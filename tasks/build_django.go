package tasks

import (
	"fmt"
	"log"

	"github.com/plexcorp-pty-ld/tuxdeploy/core"
	"github.com/plexcorp-pty-ld/tuxdeploy/scripts"
	"github.com/plexcorp-pty-ld/tuxdeploy/termio"
)

func getStepsDjango(config *core.AppConfig) []core.BuildStep {
	var steps []core.BuildStep
	steps = append(steps, core.BuildStep{
		StepName: "Setup Postgresql", Code: scripts.SetupPostgreSQL(config), RunLocal: false},
	)

	steps = append(steps, core.BuildStep{
		StepName: "Install Django APT Packages", Code: scripts.GetPythonInitialSetupScript(config), RunLocal: false},
	)

	steps = append(steps, core.BuildStep{
		StepName: "Setup Virtual Environment", Code: scripts.GetVenvSetup(config), RunLocal: false},
	)

	steps = append(steps, core.BuildStep{StepName: "Generate deploy key", Code: scripts.GenerateDeployKey(config), RunLocal: false})

	return steps
}

func BuildServerDjango(config *core.AppConfig) {

	termio.DrawTerminalHeader("Setup of Django essentials such as a virtual environment, gunicorn and nginx.")

	sshConfig := core.GetNewSSHConfig(config)
	client, err := core.GetSshClient(config, sshConfig)

	if err == nil {
		defer client.Close()
	} else {
		log.Fatal("Sorry, there was a problem connecting to your server:", err)
		return
	}

	steps := getStepsDjango(config)

	for stepNum, step := range steps {

		response, err := core.RunSshCommand(client, &step, stepNum)

		if err != nil {
			termio.PrintError("Oops, step: "+step.StepName+" has errored with message: "+err.Error()+".", 60)
			return
		}

		if err != nil {
			termio.PrintError("Oops, step: "+step.StepName+" has errored with message: "+err.Error()+". Response: "+string(response), 60)
			return
		}

		stepNum++
		fmt.Println()
		termio.PrintRegularText(">>> BUILD LOG - "+step.StepName+" <<<<", 0)
		fmt.Println(string(response))
		termio.PrintRegularText(">>> END BUILD LOG - "+step.StepName+" <<<<", 0)
		fmt.Println()

	}

	msg := termio.GetColorText("#0B832C")
	fmt.Println(msg.Render("Yay! all steps for Django environment complete."))
}
