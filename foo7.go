package main

import (
	"fmt"
)

type S struct {
	x int32
}
func (S) p2 () error {
	return nil
}
func (s *S) p () error {
	fmt.Print(s)
	return nil
}

func main() {
	s := new(S)
	(*s).p()
	(*s).p2()
}
