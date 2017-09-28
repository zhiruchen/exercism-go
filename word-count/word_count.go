package wordcount

import (
	"regexp"
	"strings"
)

const testVersion = 3

// Use this return type.
type Frequency map[string]int

// Just implement the function.
func WordCount(phrase string) Frequency {
	re := regexp.MustCompile("[a-zA-Z']+|\\d+")
	wordFreq := make(Frequency)

	words := re.FindAllString(phrase, -1)
	for _, word := range words {
		if word[0] == []byte("'")[0] && word[len(word)-1] == []byte("'")[0] {
			word = word[1 : len(word)-1]
		}
		word := strings.ToLower(word)
		wordFreq[word] = wordFreq[word] + 1
	}

	return wordFreq
}
