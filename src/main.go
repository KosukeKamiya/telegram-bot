package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"telegram-bot/src/util"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func main() {
	http.HandleFunc("/webhook/", wenhookHandler)
	http.HandleFunc("/", errorHandler)
	appengine.Main()
}

func wenhookHandler(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	if req.Method != "POST" {
		log.Errorf(ctx, "Method is not POST.")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	// get Telegram BOT token
	tgToken := strings.Replace(req.URL.Path, "/webhook/", "", 1)
	log.Infof(ctx, "token: "+tgToken)

	// Parse request JSON
	requestBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	log.Infof(ctx, "body: "+string(requestBody))

	tgUpdate := new(util.TgUpdate)
	json.Unmarshal(requestBody, tgUpdate)

	updateType := tgUpdate.UpdateType()
	log.Infof(ctx, "updateType: %d", updateType)

	// Switch by type of webhook
	switch {
	// in case of message, send message.
	case updateType == util.MessageType:
		log.Infof(ctx, "From: %d", tgUpdate.Message.Chat.ID)
		log.Infof(ctx, "Text: %s", tgUpdate.Message.Text)
		util.SendMessage(ctx, tgUpdate.Message.Chat.ID, tgUpdate.Message.Text, tgToken)
	}
}

func errorHandler(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	log.Errorf(ctx, "URL is wrong.")
	res.WriteHeader(404)
}
