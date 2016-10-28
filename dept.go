package main

import (
	"fmt"
	"github.com/fatih/structs"
)

type Dept struct {
	Name        string
	Count       int
	ParentDept  *Dept
	NextSibling *Dept
}

func main() {
	nullDept := &Dept{}

	root := &Dept{
		Name:        "root",
		Count:       0,
		ParentDept:  nil,
		NextSibling: nil,
	}

	f := structs.Fields(root)
	fmt.Println(structs.IsStruct(root))
	fmt.Println(structs.IsZero(nullDept))

	m := make(map[string]interface{})
	// 将 struct 提取出来成为一个 map
	structs.FillMap(root, m)
	fmt.Println(m)

	// 是不是有一个零值
	fmt.Println(structs.HasZero(root))

	// 使用一个 map 填充一个 struct ???
	structs.Struct{}

	for i, field := range f {
		fmt.Printf("[%d] %+v\n", i, field.Name())
	}
}
