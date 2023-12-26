package core

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/melbahja/goph"
	"github.com/plexcorp-pty-ld/tuxdeploy/termio"

	"github.com/BurntSushi/toml"
	"golang.org/x/crypto/ssh"
)

func GetGlobalAppConfigPath() string {
	globalConfigPath, _ := os.UserHomeDir()
	globalConfigPath = globalConfigPath + "/.tuxdeploy.toml"
	return globalConfigPath
}

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
	if HasPassphrase(appConf.SSHKey) {
		passphrase = termio.PromptPassword("Please enter the passphrase for this key?")
	}

	auth, err := goph.Key(appConf.SSHKey, passphrase)
	if err != nil {
		return nil, err
	}

	sshConfig.Auth = auth
	sshConfig.Callback = VerifyHost

	return goph.NewConn(sshConfig)
}

func GetTomlConfig() *AppConfig {

	_, foundLocalConfig := os.Stat("./.tuxdeploy.toml")
	globalConfigPath := GetGlobalAppConfigPath()
	_, foundGlobalConfig := os.Stat(globalConfigPath)
	var config = &AppConfig{}

	if foundLocalConfig != nil && foundGlobalConfig != nil {
		return config
	}

	filename := globalConfigPath
	if foundLocalConfig == nil {
		filename = "./.tuxdeploy.toml"
	}

	_, err := toml.DecodeFile(filename, config)
	fmt.Println(err, foundLocalConfig)
	return config
}

func GetIntialSSHConfig(config *AppConfig) *goph.Config {
	return &goph.Config{
		User: config.Username,
		Addr: config.Server,
		Port: uint(config.Port),
	}
}

func GetNewSSHConfig(config *AppConfig) *goph.Config {
	return &goph.Config{
		User: config.NewUsername,
		Addr: config.Server,
		Port: uint(config.NewSSHPort),
	}
}
