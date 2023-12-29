package core

import "regexp"

type AppConfig struct {
	Server struct {
		Address     string `toml:"address"`
		Username    string `toml:"username"`
		SSHKey      string `toml:"ssh_key"`
		Port        int    `toml:"port"`
		NewUsername string `toml:"new_username"`
		NewSSHPort  int    `toml:"new_ssh_port"`
	} `toml:"server"`
	Project struct {
		Domain          string `toml:"domain"`
		ProjectName     string `toml:"project_name"`
		ProjectGit      string `toml:"project_git"`
		GunicornWorkers int    `toml:"gunicorn_workers"`
		GunicornPort    int    `toml:"gunicorn_port"`
	} `toml:"project"`
}

func (c *AppConfig) GetProjectPath() string {
	return "/var/www/" + c.Project.Domain
}

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func (c *AppConfig) GetProjecUsername() string {
	return nonAlphanumericRegex.ReplaceAllString(c.Project.ProjectName, "")
}

type BuildStep struct {
	StepName string
	Code     string
	RunLocal bool
}
