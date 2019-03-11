package translate

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

func Translate(origin *string) (string, error) {

	fmt.Printf("orig %v", *origin)
	ctx := context.Background()

	// Creates a client.
	client, err := translate.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	target, err := language.Parse("en")
	if err != nil {
		log.Fatalf("Failed to parse target language: %v", err)
	}

	// Translates the text into Russian.
	translations, err := client.Translate(ctx, []string{(*origin)[0:500]}, target, nil)
	if err != nil {
		log.Fatalf("Failed to translate text: %v", err)
		return "", err
	}

	fmt.Printf("Translation: %v\n", translations[0].Text)
	return translations[0].Text, nil
}
