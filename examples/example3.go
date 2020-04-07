package main

import (
	"fmt"

	"github.com/zajann/finder"
)

func main() {
	text := "Lorem ipsum example of lorem ipsum."

	str := "lorem ipsum"

	n, ok := finder.Sentence(text, str)
	fmt.Println(n)
	fmt.Println(ok)
}
