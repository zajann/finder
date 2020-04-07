package main

import (
	"fmt"

	"github.com/zajann/finder"
)

func main() {
	text := `
h sdfgsdfgdfgdg http://www.dffff.com dfgdfgfdgn dfgdfgdfgdgflslfl      sfdgsgsdfg   sdfgsghfsgzajan.com https://www.totohot.net/bbs/board.php?bo_table=mt&page=1 http://www.dffff.com
	`

	urls := finder.URL(text)
	fmt.Println(urls)

}
