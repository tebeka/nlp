package nlp

import (
	"regexp"
	"strings"

	"github.com/tebeka/nlp/stemmer"
)

var (
	wordRe = regexp.MustCompile("[[:alpha:]]+")
)

// Tokenize returns a slice of tokens found in text
func Tokenize(text string) []string {
	words := wordRe.FindAllString(text, -1)
	var tokens []string
	for _, w := range words {
		// TODO: stem
		token := stemmer.Stem(strings.ToLower(w))
		if token != "" {
			tokens = append(tokens, token)
		}
	}
	return tokens
}
