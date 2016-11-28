package main

import (
	"fmt"
	"encoding/json"
	"reflect"
)

type P struct {
	x int
}
func (P) Error() string {
	return ""
}

func main() {
	x := "{}"
	data := make(map[string]interface{})

	err := json.Unmarshal([]byte(x), &data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}

	fmt.Println("125xxxx"[:len("124xxxx")-4])

	var p error
	p = P{}
	fmt.Println(reflect.TypeOf(p))

	switch v := p.(type) {
	case P:
		fmt.Println(v, "123")
	default:
		fmt.Println(v, "456")
	}
}
