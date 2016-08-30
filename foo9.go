package main

import "fmt"

type Foo struct {
	x int
}
func (f Foo) action() {}

// 根据 struct 创建的新类型 还是可以访问到原始数据结构的
// 但是绑定到原有类型上的 method 就不行辣
type Bar Foo

func main() {
	f := Foo{}
	f.action()

	b := Bar{}
	fmt.Println(b.x)
}
