package main

import "strings"

func truncateText(text string, maxLen int) string {
	text = strings.ReplaceAll(text, "\n", " ")
	if len(text) > maxLen {
		return text[:maxLen-3] + "..."
	}
	return text
}
