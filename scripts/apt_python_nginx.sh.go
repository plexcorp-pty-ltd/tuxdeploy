package scripts

import (
	"strings"

	"github.com/plexcorp-pty-ld/tuxdeploy/core"
)

func GetPythonInitialSetupScript(config *core.AppConfig) string {

	script := `
sudo apt update -y
sudo apt install python3-venv python3-dev libpq-dev nginx -y

sudo touch /etc/#project#.env
	`

	script = strings.ReplaceAll(script, "#project#", config.GetProjecUsername())
	return script
}
