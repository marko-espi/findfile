# Search File
This is a small go package that find a file based on the given extesion and return the number of matches or error if it finds one.

## How to use
Simple example on how to use the package
### Install the package
```sh
go get "github.com/marko-espi/findfile"
```
### Example:
```go
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
```