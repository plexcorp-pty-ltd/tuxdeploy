package scripts

import (
	"fmt"
	"os"
	"strings"

	"github.com/plexcorp-pty-ld/tuxdeploy/core"
)

func GetSshUserSetup(config *core.AppConfig) string {
	script :=
		`
echo "Adding user : #username#"
sudo useradd -m -s /bin/bash #username#

if ! sudo getent passwd "#username#" >/dev/null; then
    echo "Error: User #username# does not exist."
    exit 1
fi

echo '#username# ALL=(ALL) NOPASSWD: ALL' | sudo tee -a /etc/sudoers
sudo usermod -aG sudo #username#

echo "Turn off SSH password authentication"
sudo sed -i 's/.*PasswordAuthentication.*/PasswordAuthentication no/' /etc/ssh/sshd_config
sudo service ssh restart

echo "Setup authorized key for #username#"
mkdir -p /home/#username#/.ssh
touch /home/#username#/.ssh/authorized_keys

echo "#PUBKEY#" >> /home/#username#/.ssh/authorized_keys

chmod 600 /home/#username#/.ssh/authorized_keys
chmod 700 /home/#username#/.ssh
chown -R #username#:#username# /home/#username#

mkdir -p /var/www/
chown -R www-data:www-data /var/www/
	`

	script = strings.ReplaceAll(script, "#username#", config.Server.NewUsername)
	pubkey, err := os.ReadFile(config.Server.SSHKey + ".pub")
	if err != nil {
		fmt.Println(err)
	}

	script = strings.ReplaceAll(script, "#PUBKEY#", string(pubkey))

	return script
}
