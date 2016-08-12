package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.5
	FreezingC Celsius = 0
	BoiiingC Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c * 9 / 5 + 32)
}

func FToCiF(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

