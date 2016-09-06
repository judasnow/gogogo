package main

import "fmt"
import "reflect"

// slice 有两个特殊的属性 len, cap

func main() {
	var s []int
	fmt.Println(s)

	s2 := make([]int, 50, 100)

	i := 10240000

	fmt.Println(cap(s))
	fmt.Println(cap(s2))

	for j := 1; j <= i; j++ {
		s = append(s, j)
	}

	s2 = append(s2, 1)
	fmt.Println(cap(s2))

	fmt.Println(cap(s))

	// 这里加 ... 说明定义的是一个数组而不是一个 slice
	nums := [...]string{1:"one", 2:"two", 3:"three", 4:"four"}
	nums2 := []string{1:"one", 2:"two", 3:"three", 4:"four"}
	fmt.Println(reflect.TypeOf(nums))
	fmt.Println(reflect.TypeOf(nums2))

	// 这里会有一个 panic
	// nums2[5] = "five"
	// 使用 append 的话就没有问题 在底层会判断 cap 如果
	// cap 不够的话就需要添加新的空间
	nums2 = append(nums2, "five")
}
