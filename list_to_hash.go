package main

import (
	"fmt"
)

func listToMap(list interface{}) map[string]bool {
	m := make(map[string]bool)

	switch list.(type) {
	case []string:
		for _, key := range list.([]string) {
			m[key] = true
		}
	case []int:
		for _, key := range list.([]int) {
			m[string(key)] = true
		}
	default:
		fmt.Println("no fit type")
	}

	return m
}

func main() {
	fmt.Printf("map: %+v\n", listToMap([]string{"foo", "bar", "baz"}))
	fmt.Printf("map: %+v\n", listToMap([]int{1, 2, 3}))
}
