# Welcome to TuxDeploy

TuxDeploy is a CLI tool to make server management easier for Django developers. The tool will provision an Ubuntu 22.04 server within minutes and install all the essential packages you need to run any Django projects.

# Getting started

You will need Golang installed. Learn more here: [https://go.dev/doc/install](https://go.dev/doc/install) . Once Golang is installed you can run the following:
```golang
go mod tidy
go run main.go
```

## Config file

You will need to setup a **.tuxdeploy.toml** , this will contain all the server and project settings. See below an example:
```yaml
[server]
address = "192.168.1.1"
username = "root"
ssh_key = "/home/kevin/.ssh/testkey"
port = 22
new_username = "webadmin"
new_ssh_port = 2022

[project]
project_name = "mydjangoapp"
project_git = "git@github.com:plexcorp-pty-ltd/testdjango.git"
project_nginx = "nginx/{project_name}"
project_systemctl = "systemctl/{project_name}.service"
```

## Setup a blank server

TuxDeploy will only work with SSH keys because this providers a greater deal of security. Most hosting providers such as Digital Ocean, Linode and AWS allow for setting up servers with SSH keys instead of passwords.

Once you have inputted the correct **address** and **ssh_key** path (this is your private key), simply run:
```bash
go run main.go
```
This command will SSH into the server and perform essential setup tasks such as SSH hardening, install APT packages and setting up your firewall.