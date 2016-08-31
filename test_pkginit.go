package main

import _ "gogogo/test_pkginit"
import "fmt"


// main 的 init 最后执行
func init() {
	fmt.Printf("init main\n")
}

// 多个 init 函数会按顺序执行
func init() {
	fmt.Printf("init2 main\n")
}

func main() {}
