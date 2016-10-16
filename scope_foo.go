package main

import "fmt"

func main() {
	err1, err2 := 1, 2
	fmt.Println(&err1, &err2)
	// 左边存在两个变量的情况，会一起重新进行声明time
	// 这里的地址没有变，不要混淆 作用域 以及 life
	err2, err3 := 3, 4
	fmt.Println(&err1, &err2, &err3)
}
