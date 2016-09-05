package main

import "fmt"

// 两个类似的 struct 类型之间的相互赋值

type Foo struct {
	X int
	Y int
}

type Bar struct {
	X int
}

func main() {
	foo := Foo{1,2}
	
	// 直接类型转换不能
	// bar := (Bar)(foo)

	fmt.Print(bar)
}
