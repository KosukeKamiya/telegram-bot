# telegram-bot

## deploy.sh

```deploy.sh
cd `dirname $0`
cd src
rm -rf ~/.config/gcloud/logs/
gcloud app deploy app.yaml --project {appengine project id} -q --version version1
wait
```
