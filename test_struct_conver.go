package main


type Foo struct {
	Name string `json:"name" gorm:"size:32"`
}

func main() {
	foo := Foo{Name: "1024"}
}
