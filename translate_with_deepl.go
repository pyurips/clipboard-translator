package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func translateWithDeepL(url, apiKey, text, sourceLang, targetLang string) (string, string, error) {
	reqBody := DeepLRequest{
		Text:       []string{text},
		TargetLang: targetLang,
	}

	if sourceLang != "" {
		reqBody.SourceLang = sourceLang
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", "", fmt.Errorf("error creating request JSON: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", "", fmt.Errorf("error creating HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "DeepL-Auth-Key "+apiKey)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("error sending request to DeepL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", "", fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(bodyBytes))
	}

	result := DeepLResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", "", fmt.Errorf("error decoding response: %w", err)
	}

	if len(result.Translations) == 0 {
		return "", "", fmt.Errorf("no translations returned")
	}

	return result.Translations[0].Text, result.Translations[0].DetectedSourceLanguage, nil
}
