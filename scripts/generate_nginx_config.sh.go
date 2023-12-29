package scripts

import (
	"fmt"
	"strings"

	"github.com/plexcorp-pty-ld/tuxdeploy/core"
)

var nconfig = `
server {
    listen 80;
    server_name #domain#;

    location = /favicon.ico { access_log off; log_not_found off; }

	location /static/ {
        root #PROJECT_PATH#/app/static;
    }

    location / {
        include proxy_params;
        proxy_pass http://127.0.0.1:#GUNICORN_PORT#;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    error_page 500 502 503 504 /50x.html;
    location = /50x.html {
        root /usr/share/nginx/html;
    }
}

`

func GenerateNginxConfig(config *core.AppConfig) string {
	script := `echo "` + nconfig + `" | sudo tee /etc/nginx/sites-available/#domain#`
	script = strings.ReplaceAll(script, "#project#", config.Project.ProjectName)
	script = strings.ReplaceAll(script, "#domain#", config.Project.Domain)
	script = strings.ReplaceAll(script, "#username#", config.GetProjecUsername())
	script = strings.ReplaceAll(script, "#PROJECT_PATH#", config.GetProjectPath())
	script = strings.ReplaceAll(script, "#GUNICORN_PORT#", fmt.Sprintf("%d", config.Project.GunicornPort))

	return script
}
