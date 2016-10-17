package main

import "fmt"
import "unicode/utf8"

func main() {
	// len 的单位是 byte
	s1 := "琨斤拷"
	fmt.Printf("%b,%b,%b\n", s1[0], s1[1], s1[2])

	// for 的单位是 码点
	for _, x := range "琨斤拷" {
		fmt.Println(x)
	}

	// 获取 unicode 码点个数
	fmt.Println(utf8.RuneCountInString(s1))

	res := make(map[string]bool)
	fmt.Println(res["1024"])
}


// type RpcHandler struct {
// 	func handler
// 	// 如何表现一个 任意结构
// 	args interface {}
// 	reply interface{}
// }
