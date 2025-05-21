package hello

import (
	"strings"
)

type Language string

const (
	English Language = "english"
	Spanish Language = "spanish"
	French  Language = "french"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "

	englishDefaultRecipient = "World"
)

var languagePrefixes = map[Language]string{
	English: englishHelloPrefix,
	Spanish: spanishHelloPrefix,
	French:  frenchHelloPrefix,
}

func GetLanguagePrefix(language string) string {
	prefix, ok := languagePrefixes[Language(strings.ToLower(language))]
	if !ok {
		return languagePrefixes[English]
	}
	return prefix
}

func Hello(recipient, language string) string {
	if recipient == "" {
		recipient = englishDefaultRecipient
	}

	prefix := GetLanguagePrefix(language)
	return prefix + recipient
}
