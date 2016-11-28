package main

import(
	"fmt"
	"time"
	"strconv"
)

func main() {
	_now, err := time.Parse(
		"20060102",
		strconv.Itoa(20161121))

	fmt.Println(err)

	// firstDay := now.AddDate(0, 0, -1 * (6 + int(now.Weekday())))
	var weekday int
	var _weekday = int(_now.Weekday())
	if _weekday > 0 {
		weekday = _weekday - 1
	} else {
		weekday = 6
	}

	firstDay := _now.Add(time.Hour * 24 * time.Duration((7 + weekday)) * -1)
	fmt.Println(_now , weekday)
	fmt.Println(firstDay)
}
