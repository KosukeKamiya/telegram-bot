package util

import (
	"context"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

// Translation returns string that translated from language2 to language1
// TBD : language detection
func Translation(ctx context.Context, sourceText string, language1 string, language2 string) string {

	client, err := translate.NewClient(ctx)
	if err != nil {
		return ""
	}
	text := sourceText
	// Sets the target language.
	target, err := language.Parse(language1)
	if err != nil {
		return ""
	}
	translations, err := client.Translate(ctx, []string{text}, target, nil)
	if err != nil {
		return ""
	}
	return translations[0].Text
}
