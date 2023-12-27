package termio

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func DrawMainMenu() {
	DrawTerminalHeader("Tux Deploy is a powerful CLI utility that helps you build ubuntu servers fast. What would you like to do?")

	var commandText = GetColorText("#5881af")

	fmt.Println(commandText.Render("build           -- Build a new Ubuntu server from scratch."))
	fmt.Println(commandText.Render("django          -- Setup Gunicorn + Nginx + virtualenv + PostgreSQL + Redis."))
	// fmt.Println(commandText.Render("dj-setup        -- Clone django project on server and setup configs."))
	// fmt.Println(commandText.Render("deploy          -- deploy current project."))
	// fmt.Println(commandText.Render("fwa             -- Add firewall rule."))
	// fmt.Println(commandText.Render("fwl             -- list all firewall rules."))
	// fmt.Println(commandText.Render("top            -- Show snapshot of running processes."))
	// fmt.Println(commandText.Render("ptop            -- See whats running on a particular port. --ptop 8000"))
	// fmt.Println(commandText.Render("nstat           -- show networking stats."))
	// fmt.Println(commandText.Render("cat             -- cat a file. --cat /var/log/ngix/error.log"))
	// fmt.Println(commandText.Render("cp              -- copy a file to server. --cp files.zip /tmp/"))
	// fmt.Println(commandText.Render("rcp             -- copy a file from server. --cp /tmp/test.sql ."))
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
