package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

const (
	DEEPL_API_KEY = "DEEPL_API_KEY"
)

var allowedKeys = []string{DEEPL_API_KEY}

func LoadEnvVariables() {
	workDir, err := os.Getwd()
	if err != nil {
		panic("[ENV] Error getting working directory: " + err.Error())
	}

	filePath := filepath.Join(workDir, ".env")

	file, err := os.Open(filePath)
	if err != nil {
		panic("[ENV] Error opening .env file: " + err.Error())
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		panic("[ENV] Error getting .env file info: " + err.Error())
	}

	if fileInfo.Size() == 0 {
		panic("[ENV] The .env file is empty")
	}

	foundKeys := make(map[string]bool)
	for _, key := range allowedKeys {
		foundKeys[key] = false
	}

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		panic("[ENV] Error creating scanner for .env file: " + err.Error())
	}

	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			panic(fmt.Sprintf("[ENV] Invalid format at line %d: %s", lineNum, line))
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if len(value) > 1 && (strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"")) ||
			(strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'")) {
			value = value[1 : len(value)-1]
		}

		isAllowed := slices.Contains(allowedKeys, key)

		if !isAllowed {
			panic(fmt.Sprintf("[ENV] Key '%s' is not allowed. Allowed keys: %s",
				key, strings.Join(allowedKeys[:], ", ")))
		}

		if err := os.Setenv(key, value); err != nil {
			panic("[ENV] Error setting environment variable " + key + ": " + err.Error())
		}

		if _, exists := foundKeys[key]; exists {
			foundKeys[key] = true
		}
	}

	if err := scanner.Err(); err != nil {
		panic("[ENV] Error reading .env file: " + err.Error())
	}

	var missingKeys []string
	for key, found := range foundKeys {
		if !found {
			missingKeys = append(missingKeys, key)
		}
	}

	if len(missingKeys) > 0 {
		panic(fmt.Sprintf("[ENV] Required environment variables missing: %s",
			strings.Join(missingKeys, ", ")))
	}
}
