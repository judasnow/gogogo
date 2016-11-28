package main


import (
	"fmt"
	"math"
	"database/sql"
)

// 考虑值范围的情况下 将 int64 转换为 int 类型
func Int64ToInt(x int64) int {
	var y int
	if x >= math.MaxInt32 {
		y = math.MaxInt32
		return y
	} else if x <= math.MinInt32 {
		y = math.MinInt32
		return y
	} else {
		y = int(x)
		return y
	}
}

func main () {
	var s sql.NullString
	s.Scan("123")
	if s.Valid {
		fmt.Println(s.String)
	} else {
		fmt.Println(s.Valid)
	}

	var x int64 = math.MaxInt64
	fmt.Println(Int64ToInt(x))

	x = math.MinInt64
	fmt.Println(Int64ToInt(x))

	x = 3
	fmt.Println(Int64ToInt(x))
}
