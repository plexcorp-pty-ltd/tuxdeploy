package termio

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var TUX_DEPLOY_LOGO = `⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⠤⠤⢤⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⣸⡇⠀⠘⣳⣗⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⢀⠔⠉⠀⠈⠉⠉⠁⠀⠉⠢⡀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⡰⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢢⡀⠀⠀⠀⠀
⠀⠀⢀⣜⠠⠤⡄⠐⠒⠂⠖⠀⠀⠒⢲⠒⠂⠤⣵⡀⠀⠀⠀
⠀⠀⠋⡄⠀⠀⢡⠀⠀⠀⢸⠀⠀⠀⢸⠀⠀⠀⡇⠸⠀⠀⠀
⠀⠰⠀⠃⢀⣀⡸⠤⠤⠤⠼⠤⠀⠀⠂⠃⠠⠤⠼⢄⡇⠀⠀
⠀⠳⡖⠉⠁⡇⠀⠀⠀⠀⢠⡶⠆⠀⠀⠀⠰⣦⠀⢸⠀⠀⠀
⠀⠀⢣⠀⠀⠃⠀⠀⠀⠀⠈⠉⠁⠰⣨⡥⠈⠉⠀⠀⢧⠀⠀
⠀⠀⠈⣆⠀⠈⢄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡜⠀⠀
⠀⠀⠀⣸⠦⠤⢄⡡ TuxDeploy⢀⡴
`
var TOML_CONFIG_EXAMPLE = `
server = "192.168.1.1"   
username = "root"         
ssh_key = "/home/kevin/.ssh/id_rsa" 
port = 22       
new_username = "webadmin" 
new_ssh_port = 9022
`

func PrintError(msgTxt string, Width int) {
	msg := GetErrorText()
	if Width != 0 {
		fmt.Println(msg.Width(Width).Render(msgTxt))
	} else {
		fmt.Println(msg.Render(msgTxt))
	}
}

func PrintRegularText(msgTxt string, Width int) {
	msg := GetColorText("#dcebfc")
	if Width != 0 {
		fmt.Println(msg.Width(Width).Render(msgTxt))
	} else {
		fmt.Println(msg.Render(msgTxt))
	}
}

func GetErrorText() lipgloss.Style {
	return lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF3333"))
}
func GetColorText(color string) lipgloss.Style {
	return lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(color))
}

func PrintAppConfigNotFound(globalConfigPath string) {
	PrintError("Please create a ./tuxdeploy.toml file either in the current directory or "+globalConfigPath+". Your file should look like the following:\n", 60)

	PrintRegularText(TOML_CONFIG_EXAMPLE, 0)
	PrintRegularText("The \"new_\" configs are used to secure your server. SSH access will be disabled for the root user, therefore once the setup is complete - you can only access your server via the \"new_username\" user and the \"new_port\" SSH port.", 60)

}