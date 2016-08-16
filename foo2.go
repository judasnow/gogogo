package main

import "fmt"
import "unsafe"

type User struct {
	id int64
	name string
}

func main() {
	var v User
	v.id = 1
	v.name = "v"

	var n int32 = 1
	var s string = "v"

	fmt.Println(unsafe.Sizeof(n))
	fmt.Println(unsafe.Sizeof(s))
	fmt.Println(*&v.id)
}
