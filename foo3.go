package main

import "fmt"
import "unsafe"

// 结构体

type T struct{ a, b int }

type Foo struct {
    foo int
    bar string
}

type Point struct {
    x int
    y int
}

// 匿名成员
type Circle struct {
    Point
    Radius int
}

type Wheel struct {
    Circle
    Spokes int
}

func foobar(foo Foo) {
    foo.foo = 1
    fmt.Printf("%x\n", &foo)
    fmt.Println(foo.foo)
}

func main() {
    // 可以部分的进行初始化 没有必要
    foo := Foo{bar: "a"}
    fmt.Printf("%x\n", &foo)
    fmt.Println(foo.foo)
    foobar(foo)
    fmt.Printf("%x\n", &foo)
    fmt.Println(foo.foo)
    fmt.Println(foo.foo)
    // new(Foo)

    // 没有给 Wheel 分配空间??????
    // 仅仅只是一个指针 因此不需要make
    var w Wheel
    fmt.Println(unsafe.Sizeof(w))
    w.x = 1
    fmt.Println(unsafe.Sizeof(w))

    // 直接写的话 并不会出现 ch4 中描述的错误
    // 猜想应该是 包的问题
    t1 := T{a: 1, b: 2}
    t2 := T{1, 2}

    fmt.Println(t1, t2)
}
