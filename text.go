package finder

import "strings"

type WordCount map[string]int

func findWords(str string, words []string) (WordCount, bool) {
	found := false
	wordList := strings.Fields(str)
	wordMap := make(WordCount)
	for _, w := range wordList {
		if _, ok := wordMap[w]; !ok {
			wordMap[w] = 1
		} else {
			wordMap[w] += 1
		}
	}

	result := make(WordCount)
	for _, w := range words {
		if n, ok := wordMap[w]; ok {
			result[w] = n
			found = true
		}
	}
	return result, found
}

func findSentence(src string, str string) (int, bool) {
	found := false
	cnt := strings.Count(src, str)
	if cnt > 0 {
		found = true
	}
	return cnt, found
}
