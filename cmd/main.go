package main

import (
	"fmt"
	"os"

	"github.com/marko-espi/findfile"
)

func main() {
	fs := os.DirFS("/Users/marcoespinoza/src")
	count, err := findfile.File(fs, ".go")
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
}
