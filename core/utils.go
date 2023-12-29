package core

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"

	"github.com/adhocore/chin"
	"github.com/melbahja/goph"
	"github.com/plexcorp-pty-ld/tuxdeploy/termio"

	"github.com/BurntSushi/toml"
	"golang.org/x/crypto/ssh"
)

func HasPassphrase(filePath string) bool {
	keyBytes, err := os.ReadFile(filePath)
	if err != nil {
		return false
	}

	_, err = ssh.ParsePrivateKey(keyBytes)
	if err != nil && strings.Contains(err.Error(), "passphrase") {
		return true
	}

	return false
}

func VerifyHost(host string, remote net.Addr, key ssh.PublicKey) error {
	return nil
}

func GetSshClient(appConf *AppConfig, sshConfig *goph.Config) (*goph.Client, error) {
	passphrase := ""
	if HasPassphrase(appConf.Server.SSHKey) {
		passphrase = termio.PromptPassword("Please enter the passphrase for this key?")
	}

	auth, err := goph.Key(appConf.Server.SSHKey, passphrase)
	if err != nil {
		return nil, err
	}

	sshConfig.Auth = auth
	sshConfig.Callback = VerifyHost

	return goph.NewConn(sshConfig)
}

func GetTomlConfig() *AppConfig {

	filename := "./.tuxdeploy/config.toml"
	_, foundConfig := os.Stat(filename)
	var config = &AppConfig{}

	if foundConfig != nil {
		return config
	}

	_, err := toml.DecodeFile(filename, config)
	if err != nil {
		fmt.Println("Error reading TOML configuration file: ", err)
	}
	return config
}

func GetIntialSSHConfig(config *AppConfig) *goph.Config {
	return &goph.Config{
		User: config.Server.Username,
		Addr: config.Server.Address,
		Port: uint(config.Server.Port),
	}
}

func GetNewSSHConfig(config *AppConfig) *goph.Config {
	return &goph.Config{
		User: config.Server.NewUsername,
		Addr: config.Server.Address,
		Port: uint(config.Server.NewSSHPort),
	}
}

func RunSshCommand(client *goph.Client, step *BuildStep, stepNum int) ([]byte, error) {
	spinner := chin.New()
	go spinner.Start()
	termio.PrintRegularText(fmt.Sprintf("Running step[%d]: %s", stepNum+1, step.StepName), 0)

	var response []byte
	var err error
	if step.RunLocal {
		response, err = exec.Command(step.Code).Output()
	} else {
		response, err = client.Run(step.Code)

	}

	spinner.Stop()

	return response, err
}
