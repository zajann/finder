package finder

func URL(str string) []string {
	return findURL(str)
}

func Words(str string, words []string) (WordCount, bool) {
	return findWords(str, words)
}

func Sentence(src string, str string) (int, bool) {
	return findSentence(src, str)
}
