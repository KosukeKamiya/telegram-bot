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

	// Language detection
	languages, err := client.DetectLanguage(ctx, []string{sourceText})
	if err != nil {
		return ""
	}
	inputlanguage := languages[0][0].Language.String()

	target := ""
	if inputlanguage == language1 {
		target = language2
	} else if inputlanguage == language2 {
		target = language1
	} else {
		return ""
	}

	targetlanguage, err := language.Parse(target)
	if err != nil {
		return ""
	}

	translations, err := client.Translate(ctx, []string{sourceText}, targetlanguage, nil)
	if err != nil {
		return ""
	}
	return translations[0].Text
}
