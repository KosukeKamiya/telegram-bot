package util

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

const tgAPIBaseURL = "https://api.telegram.org/bot%s/sendMessage"

// SendMessage function sends telegram message.
func SendMessage(ctx context.Context, destination int64, message string, tgToken string) {
	log.Infof(ctx, "destination: %d", destination)
	log.Infof(ctx, "message: %s", message)
	log.Infof(ctx, "tgToken: %s", tgToken)

	url := fmt.Sprintf(tgAPIBaseURL, tgToken)
	requestBody := fmt.Sprintf(`{"chat_id":%d ,"text":"%s"}`, destination, message)

	client := urlfetch.Client(ctx)
	resp, err := client.Post(url, "application/json", strings.NewReader(requestBody))
	if err != nil {
		// Do something
		return
	}
	log.Infof(ctx, "HTTP POST returned status %v", resp.Status)
}
