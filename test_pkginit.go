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


// link: http://stackoverflow.com/questions/24790175/when-is-the-init-function-in-go-golang-run
var WhatIsThe = AnswerToLife()

func AnswerToLife() int {
	return 42
}

func init() {
	WhatIsThe = 0
}

func main() {
	// init 会在 main 之前被调用
	if WhatIsThe == 0 {
		fmt.Println("It's all a lie.")
	} else {
		fmt.Println("It's 42")
	}
}

