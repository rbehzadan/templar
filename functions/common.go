package functions

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Title(s string) string {
	caser := cases.Title(language.AmericanEnglish)
	return caser.String(s)
}

func Split(sep, s string) []string {
	return strings.Split(s, sep)
}
