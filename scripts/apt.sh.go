package scripts

var INIT_SERVER = `
echo 'debconf debconf/frontend select Noninteractive' | debconf-set-selections

sudo apt-get update -y
sudo apt-get upgrade -y

sudo apt-get install software-properties-common curl zip unzip git ufw libsqlite3-0 libsqlite3-dev python3-pip -y
`
