package termio

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func DrawMainMenu() {
	DrawTerminalHeader("Tux Deploy is a powerful CLI utility that helps you build ubuntu servers fast. What would you like to do?")

	var commandText = GetColorText("#5881af")

	fmt.Println(commandText.Render("build           -- Build a new Ubuntu server from scratch."))
	fmt.Println(commandText.Render("build-dj        -- Setup Gunicorn + Nginx + virtualenv."))
	fmt.Println(commandText.Render("deploy          -- deploy current project."))
	fmt.Println(commandText.Render("pg              -- Setup a Postgresql server."))
	fmt.Println(commandText.Render("msq             -- Setup a MySQL server."))
	fmt.Println(commandText.Render("fwa             -- Add firewall rule."))
	fmt.Println(commandText.Render("fwd             -- delete firewall rule."))
	fmt.Println(commandText.Render("fwl             -- list all firewall rules."))
	fmt.Println(commandText.Render("cl              -- list all crons."))
	fmt.Println(commandText.Render("ca              -- Add cron JOB."))
	fmt.Println(commandText.Render("cad             -- Delete cron JOB."))
	fmt.Println(commandText.Render("top            -- Show snapshot of running processes."))
	fmt.Println(commandText.Render("ptop            -- See whats running on a particular port. --ptop 8000"))
	fmt.Println(commandText.Render("nstat           -- show networking stats."))
	fmt.Println(commandText.Render("cat             -- cat a file. --cat /var/log/ngix/error.log"))
	fmt.Println(commandText.Render("cp              -- copy a file to server. --cp files.zip /tmp/"))
	fmt.Println(commandText.Render("rcp             -- copy a file from server. --cp /tmp/test.sql ."))
}

func DrawTerminalHeader(msg string) {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#2E97C3"))

	fmt.Println(style.Render(TUX_DEPLOY_LOGO))

	var helptext = lipgloss.NewStyle().Bold(true).
		Foreground(lipgloss.Color("#dcebfc")).Width(60).PaddingBottom(1)
	fmt.Println(helptext.Render(msg))
}
