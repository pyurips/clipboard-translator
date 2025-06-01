package main

const (
	MAX_CHAR_LIMIT        = 100
	API_URL               = "https://api-free.deepl.com/v2/translate"
	TRANSLATION_FROM_LANG = "EN"
	TRANSLATION_TO_LANG   = "PT-BR"
)

type DeepLRequest struct {
	Text       []string `json:"text"`
	SourceLang string   `json:"source_lang,omitempty"`
	TargetLang string   `json:"target_lang"`
}

type DeepLResponse struct {
	Translations []struct {
		DetectedSourceLanguage string `json:"detected_source_language"`
		Text                   string `json:"text"`
	} `json:"translations"`
}
