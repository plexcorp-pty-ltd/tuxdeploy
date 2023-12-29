package scripts

var deploy = `

cd #PROJECT_PATH# && sudo -u #username# git clone #GIT_URL# app_new_version

if [ -d "#PROJECT_PATH#/app" ]; then
	mkdir -p #PROJECT_PATH#/backups
    mv #PROJECT_PATH#/app #PROJECT_PATH#/backups/
fi

mv #PROJECT_PATH#/app_new_version #PROJECT_PATH#/app

`
