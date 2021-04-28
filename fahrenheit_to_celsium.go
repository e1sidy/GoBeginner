package main

import (
	"fmt"
)

type (
	Celsius    float32
	Fahrenheit float32
)

const (
	cToF = Celsius(9.0 / 5.0)
	fToC = Fahrenheit(5.0 / 9.0)
)

func toFahrenheit(t Celsius) Fahrenheit {
	return Fahrenheit((cToF * t) + 32)
}

func main() {
	var tempCelsius Celsius = 100
	// Assignment operator := determine the data type of the variable by type inference
	tempFhar := toFahrenheit(tempCelsius)
	fmt.Printf("%.2f'C is equal to %.2f'F\n", tempCelsius, tempFhar)
}
