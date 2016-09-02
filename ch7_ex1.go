package main

import (
	"fmt"
	"bufio"
)

// 练习 7.1： 使用来自ByteCounter的思路，实现一个针对对单词和行数的计数器。你会发现bufio.ScanWords非常的有用。
// 这个思路就是 由于 Fprint 的第一个参数需要传递的是一个 io.Writer 接口
// 因此可以实现这个接口的时候 进行一些特殊的处理

type wordCount int

func (wc *wordCount) Write(p []byte) (int, error) {
	a, t, err := bufio.ScanWords(p, true)
	fmt.Println(a, t, err)
	*wc += wordCount(a)
	return 1024, nil
}

func main() {
	var wc wordCount = 0

	// fmt.Fprintf(wc, "abcd", "2048")
	wc.Write([]byte("abcd"))
	fmt.Print(wc)
}
