package main

import (
	"fmt"
)

type ByteCounter int

// 这里为什么不能直接使用 int 类型 ???
// 也就是接收器为什么不能是 int 类型 ???
// 是因为不能直接给包外的类型定义方法
func (c ByteCounter) Write(p []byte) (int, error) {
	// 修改 ByteCounter 本身的值
    c += ByteCounter(len(p))
    return len(p), nil
}

func main() {
    var c ByteCounter

    c.Write([]byte("hello"))
    fmt.Println(c)
    c = 0
    var name = "Dolly"
    fmt.Fprintf(c, "hello, %s", name)
    fmt.Println(c)
}
