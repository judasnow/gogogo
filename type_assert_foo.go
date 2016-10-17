package main

import (
    "log"
)

// asserts that x is not nil and that the value stored in x is of type T. The notation x.(T) is called a type assertion.
// More precisely, if T is not an interface type, x.(T) asserts that the dynamic type of x is identical to the type T. In this case,
// T must implement the (interface) type of x; otherwise the type assertion is invalid since it is not possible for x to store a value of type T.
// If T is an interface type, x.(T) asserts that the dynamic type of x implements the interface T.

// 怎么理解？接口类型变量的值部分存放的是指针？

func main() {
    n := New()

    // foo := (n).(*Faz)
    // log.Println(foo)

	// 这里天真的不能加 *， 因为 n 的类型是 interface 的
    foo := (n).(*Faz)
    log.Println(foo)
}

type Foo interface {
    Bar() string
}

// 返回类型是一个接口 但是值却是一个 指针
func New() Foo {
    return &Faz{}
}

type Faz struct {
}

func (f *Faz) Bar() string {
    return `Bar`
}
