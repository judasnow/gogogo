package main

type Foo struct {
	X int
}

type Bar struct {
	Foo
}

func main() {
	b := Bar{Foo.X: 3}
	b.X = 2
}
