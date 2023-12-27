package scripts

import (
	"strings"

	"github.com/plexcorp-pty-ld/tuxdeploy/core"
)

var deploy_key_code = `
sudo mkdir -p  #project#/.ssh

sudo chown -R #username#:#username# #project#/.ssh
sudo chmod  0700 #project#/.ssh

sudo -u #username# ssh-keygen -b 2048 -t rsa -f #project#/.ssh/id_rsa -q -N ""
` + "\n" +
	`echo "Please copy the following key to your projects deploy keys:"
` +
	"echo \n" +
	`

sudo -u #username# cat #project#/.ssh/id_rsa.pub
` + "echo \n"

func GenerateDeployKey(config *core.AppConfig) string {

	script := strings.ReplaceAll(deploy_key_code, "#project#", config.GetProjectPath())
	script = strings.ReplaceAll(script, "#username#", config.GetProjecUsername())
	return script

}
