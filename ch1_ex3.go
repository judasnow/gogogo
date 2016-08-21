package main

import (
	"fmt"
	_ "strings"
	"time"
	"strings"
)

// potentially inefficient
func join1(strs []string, spliter string) string {
	var s string
	for _, str := range strs {
		s += str
		s += spliter
	}

	return s
}

func main() {
	test_strs := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	start := time.Now()
	join1(test_strs, ",")
	secs := time.Since(start).Seconds()
	fmt.Printf("%.10fs\n", secs)

	start2 := time.Now()
	strings.Join(test_strs, ",")
	secs2 := time.Since(start2).Seconds()
	fmt.Printf("%.10fs\n", secs2)
}
