package main

import (
	"os"
	_ "fmt"
)

func main() {
	for index, arg := range os.Args {
		println(index, arg)
	}
}
