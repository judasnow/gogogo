package main

import "fmt"
import "unsafe"

// 类型别名不会继承方法啊 reciver 类型不同 也是可以理解的
type Foo struct {
	Name string
}
func (Foo) Bar() {
	fmt.Println("1024")
}

type Baz Foo


type Bar struct {
	Name []Baz
}

func main() {
	// b := Baz{}
	// b.Bar()

	var b []Baz
	fmt.Printf("%+v\n", b)
	fmt.Printf("%d\n", unsafe.Sizeof(b))

	c := make([]Baz, 0)
	fmt.Printf("%+v\n", c)
}
