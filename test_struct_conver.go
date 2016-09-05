package main


// 允许多个 tag 但需要使用空格进行分割
type Foo struct {
	Name string `json:"name" gorm:"size:32"`
}

func main() {
	foo := Foo{Name: "1024"}
}
