package core

type AppConfig struct {
	Server      string `toml:"server"`
	Username    string `toml:"username"`
	SSHKey      string `toml:"ssh_key"`
	Port        int    `toml:"port"`
	NewUsername string `toml:"new_username"`
	NewSSHPort  int    `toml:"new_ssh_port"`
}
