package main

import "strings"

func cleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}

	split := strings.Fields(text)
	lower := make([]string, len(split))

	for i, word := range split {
		lower[i] = strings.ToLower(word)
	}
	return lower
}