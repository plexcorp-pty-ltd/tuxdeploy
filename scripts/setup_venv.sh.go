package scripts

import (
	"strings"

	"github.com/plexcorp-pty-ld/tuxdeploy/core"
)

var SETUP_VENV = `
  if ! sudo getent passwd "#USERNAME#" >/dev/null; then
  	sudo useradd -m -s /bin/bash #USERNAME#
  fi

  sudo mkdir -p #PROJECT_PATH#
  sudo chown -R #USERNAME#:www-data #PROJECT_PATH#
  cd #PROJECT_PATH# && sudo -u #USERNAME# python3 -m venv .venv
  sudo -u #USERNAME# #PROJECT_PATH#/.venv/bin/pip3 install gunicorn
  sudo -u #USERNAME# #PROJECT_PATH#/.venv/bin/pip3 install psycopg2
`

func GetVenvSetup(config *core.AppConfig) string {

	script := SETUP_VENV
	username := config.GetProjecUsername()
	projectPath := config.GetProjectPath()

	script = strings.ReplaceAll(script, "#PROJECT_PATH#", projectPath)
	script = strings.ReplaceAll(script, "#USERNAME#", username)

	return script
}
