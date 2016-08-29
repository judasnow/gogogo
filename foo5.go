package main

import "fmt"

func printHELLO() {
	fmt.Print("1024\n")
}

func main() {
	fmt.Print("1\n")
	defer printHELLO()
	fmt.Print("2\n")
}
