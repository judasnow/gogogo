package main

import (
	"fmt"
)

type Rpc interface {
	Call(cmd string)
}

type ClubRpc struct {

}
func (ClubRpc) Call(cmd string) {
	fmt.Println("cmd:", cmd)
}
func (ClubRpc) Call2(cmd string) {
	fmt.Println("cmd2: ", cmd)
}

func main() {
	var r *ClubRpc
	fmt.Println(r == nil)

	//var r1 ClubRpc
	 //直接声明的结构题不能和 nil 比较，因为他不是 nil 而是每一个字段都是零值的结构体类型
	//fmt.Println(r1 == nil)

	var r1 Rpc
	fmt.Println(r1 == nil)

	// 一种比较迷惑的情况就是 接口类型的变量包含一个值 但是这个值是 nil
	// 这时候 他就不是 nil 辣
	var r2 Rpc
	r2 = r
	fmt.Println(r2 == nil)
	fmt.Println(r2 == nil)
	fmt.Printf("%T,%+v\n", r2, r2)

	// 什么是类型断言，为何需要这种机制？
	// 比如这里 r2 是 interface 类型的变量 其会遮蔽原有类型的额外方法
	// 我理解的就是 针对于 接口类型的类型转换
	//r2.Call2("")

	// 比如这样强转 是不允许的，会提醒 需要类型断言
	//r3 := (*ClubRpc)(r2)
	//fmt.Printf("%T\n", r3)

	r4 := r2.(*ClubRpc)
	fmt.Printf("%T\n", r4)
}
