package main

import (
	"fmt"

	"github.com/zajann/finder"
)

func main() {
	text := "apple banana kakao phone apple phone water food food web"
	words := []string{
		"apple",
		"banana",
		"food",
	}

	count, ok := finder.Words(text, words)
	fmt.Println(ok)
	fmt.Println(count)
}
