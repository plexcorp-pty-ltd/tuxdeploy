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
[server]
address = "192.168.1.1"   
username = "root"         
ssh_key = "/home/kevin/.ssh/testkey" 
port = 22      
new_username = "webadmin" 
new_ssh_port = 2022

[project]
project_name = "mydjangoapp"
project_git = "git@github.com:plexcorp-pty-ltd/testdjango.git"
project_nginx = "nginx/{project_name}"
project_systemctl = "systemctl/{project_name}.service"
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
	msg := GetColorText(TEXT_LIGHT, TEXT_DARK)
	if Width != 0 {
		fmt.Println(msg.Width(Width).Render(msgTxt))
	} else {
		fmt.Println(msg.Render(msgTxt))
	}
}

func GetErrorText() lipgloss.Style {
	return GetColorText(ERROR_TEXT, ERROR_TEXT)
}
func GetColorText(colorLight string, colorDark string) lipgloss.Style {
	return lipgloss.NewStyle().Bold(true).Foreground(lipgloss.AdaptiveColor{Light: colorLight, Dark: colorDark})
}

func PrintAppConfigNotFound(globalConfigPath string) {
	PrintError("Please create a ./tuxdeploy.toml file either in the current directory or "+globalConfigPath+". Your file should look like the following:\n", 60)

	PrintRegularText(TOML_CONFIG_EXAMPLE, 0)
	PrintRegularText("The \"new_\" configs are used to secure your server. SSH access will be disabled for the root user, therefore once the setup is complete - you can only access your server via the \"new_username\" user and the \"new_port\" SSH port.", 60)

}
