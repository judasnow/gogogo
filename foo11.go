package main


type IFoo interface {
	Foo()
}

type Bar struct {}

func (Bar) Foo () {}

func main() {
	var i IFoo
	var j IFoo

	i = Bar{}
	j = Bar{}

	println(i)
	println(j)
}
