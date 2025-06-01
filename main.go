package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/atotto/clipboard"
)

func main() {
	LoadEnvVariables()

	apiKey := os.Getenv("DEEPL_API_KEY")

	fmt.Println("Translator started. Listening for clipboard changes...")
	fmt.Println("Press Ctrl+C to exit.")
	fmt.Println(strings.Repeat("-", 80))

	lastText := ""
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		text, err := clipboard.ReadAll()
		if err != nil {
			continue
		}

		if text == "" || text == lastText {
			continue
		}

		lastText = text

		translationText := text
		if len(translationText) > MAX_CHAR_LIMIT {
			translationText = translationText[:MAX_CHAR_LIMIT]
		}

		translation, _, err := translateWithDeepL(API_URL, apiKey, translationText, TRANSLATION_FROM_LANG, TRANSLATION_TO_LANG)

		timestamp := time.Now().Format("15:04:05")

		if err != nil {
			fmt.Printf("\n[%s] Error: %v\n", timestamp, err)
			continue
		}

		fmt.Println(strings.Repeat("-", 80))
		fmt.Println("At: ", timestamp)
		fmt.Printf("Original: %s\n", truncateText(text, 60))
		fmt.Printf("Translation: %s\n", translation)
		fmt.Println(strings.Repeat("-", 80))
	}
}
