package scripts

import (
	"fmt"
	"strings"

	"github.com/plexcorp-pty-ld/tuxdeploy/core"
)

var gconfig = `
[Unit]
Description=Gunicorn Service for #project#
After=network.target

[Service]
User=#username#
Group=www-data
WorkingDirectory=#PROJECT_PATH#/app
EnvironmentFile=/etc/#project#.env
ExecStart=#PROJECT_PATH#/.venv/bin/gunicorn -w #GUNICORN_WORKERS# -b 0.0.0.0:#GUNICORN_PORT# #project#.wsgi
StandardOutput=/var/log/#domain#.log
StandardError=/var/log/#domain#_error.log

[Install]
WantedBy=multi-user.target
`

func GenerateGunicornConfig(config *core.AppConfig) string {
	script := `echo "` + gconfig + `" | sudo tee /etc/systemd/system/#domain#.service`
	script = strings.ReplaceAll(script, "#project#", config.Project.ProjectName)
	script = strings.ReplaceAll(script, "#domain#", config.Project.Domain)
	script = strings.ReplaceAll(script, "#username#", config.GetProjecUsername())
	script = strings.ReplaceAll(script, "#PROJECT_PATH#", config.GetProjectPath())
	script = strings.ReplaceAll(script, "#GUNICORN_PORT#", fmt.Sprintf("%d", config.Project.GunicornPort))
	script = strings.ReplaceAll(script, "#GUNICORN_WORKERS#", fmt.Sprintf("%d", config.Project.GunicornWorkers))
	return script
}
