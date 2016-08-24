package main

import (
    "fmt"
	"math"
    "math/rand"
)

type Point struct {
	X float64
	Y float64
}

func (r Point) Q(q Point) float64 {
    return 1.024
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func main() {
    p := Point{1, 2}
    q := Point{3, 4}

    fmt.Println(p.Distance(q))
}
