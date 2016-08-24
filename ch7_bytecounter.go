package main

import (
	"fmt"
)

type ByteCounter int

// 这里为什么不能直接使用 int 类型 ???
// 也就是接收器为什么不能是 int 类型 ???
func (c *ByteCounter) Write(p []byte) (int, error) {
    *c += ByteCounter(len(p))
    return len(p), nil
}

func main() {
    var c ByteCounter

    c.Write([]byte("hello"))
    fmt.Println(c)
    c = 0
    var name = "Dolly"
    fmt.Fprintf(&c, "hello, %s", name)
    fmt.Println(c)
}
