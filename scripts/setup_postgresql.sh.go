package scripts

import (
	"strings"

	"github.com/plexcorp-pty-ld/tuxdeploy/core"
	"github.com/plexcorp-pty-ld/tuxdeploy/termio"
)

var setupScript = `
sudo apt install postgresql postgresql-contrib redis-server -y
cd /var/lib/postgresql

sudo -u postgres psql -c "CREATE DATABASE #project#"
sudo -u postgres psql -c "CREATE USER #project# WITH PASSWORD '#password#'"

sudo -u postgres psql -c "ALTER ROLE #project# SET client_encoding TO 'utf8'"
sudo -u postgres psql -c "ALTER ROLE #project# SET default_transaction_isolation TO 'read committed'"

sudo -u postgres psql -c "GRANT ALL PRIVILEGES ON DATABASE #project# TO #project#"

echo 'PG_DB=#project#' | sudo tee -a /etc/#project#.env
echo 'PG_DB_USER=#project#' | sudo tee -a /etc/#project#.env
echo 'PG_DB_PASSWORD="#password#"' | sudo tee -a /etc/#project#.env

`

func SetupPostgreSQL(config *core.AppConfig) string {

	var password string
	for {
		password = termio.PromptPassword("Please enter a password for this projects PostgreSQL DB?")
		passwordAgain := termio.PromptPassword("Please confirm the password entered above?")

		if password != passwordAgain {
			termio.PrintError("Password and confirm password do not match. Please try again", 60)
			continue
		}

		break
	}

	script := setupScript
	script = strings.ReplaceAll(script, "#project#", config.GetProjecUsername())
	script = strings.ReplaceAll(script, "#password#", password)

	return script
}
