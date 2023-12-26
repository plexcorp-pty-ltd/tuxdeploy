package scripts

import (
	"fmt"
	"strings"

	"github.com/plexcorp-pty-ld/tuxdeploy/core"
)

var SECURITY_POST_STEPS = `
sudo sed -i 's/^#Port #SSH_PORT#/Port #NEW_SSH_PORT#/' /etc/ssh/sshd_config
sudo sed -i 's/^#Port 22/Port #NEW_SSH_PORT#/' /etc/ssh/sshd_config

sudo apt-get install fail2ban -y
sudo cp /etc/fail2ban/jail.conf /etc/fail2ban/jail.local
sudo sed -i 's/# ignoreip = 127.0.0.1/ignoreip = 127.0.0.1/; s/# bantime = 10m/bantime = 1h/; s/# findtime = 10m/findtime = 10m/; s/# maxretry = 5/maxretry = 3/' /etc/fail2ban/jail.local
sudo service fail2ban restart

sudo sh -c "echo 'AllowUsers #username#' >> /etc/ssh/sshd_config"

sudo service ssh restart
sudo ufw allow #NEW_SSH_PORT#
sudo ufw allow 53
echo "y" | sudo ufw enable
`

func GetSecurityPostSteps(config *core.AppConfig) string {
	script := SECURITY_POST_STEPS
	script = strings.ReplaceAll(script, "#SSH_PORT#", fmt.Sprintf("%d", config.Port))
	script = strings.ReplaceAll(script, "#NEW_SSH_PORT#", fmt.Sprintf("%d", config.NewSSHPort))
	script = strings.ReplaceAll(script, "#username#", config.NewUsername)

	return script
}
