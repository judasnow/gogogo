package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := int(32)
	m := new(map[string]int)
	// 指针是即时求的一个值 并没有对应的变量存在
	// sizeof 返回的是占用的量 而并不是说明真的
	// 已经分配了多少内存
	fmt.Printf("a: %T, %d\n", &a, unsafe.Sizeof(&a))
	fmt.Printf("m: %T, %d\n", m, unsafe.Sizeof(m))
}
