package main

import "net/http"

// 被 defer 标记的函数 会在上级函数全部执行完毕之后
// 执行
func title(url string) error {
	resp, err := http.Get(url)
}

func main() {

}
