package main

import (
	"fmt"
	"vasquezruiz.com/conversions"
)

func main() {
	fmt.Println("**Conversiones**")
	bodyTemperature := conversions.Celsius(37)

	fmt.Println("La temperatura corporal promedio es", bodyTemperature)

	equiv := bodyTemperature.CalculateEquivalents()

	fmt.Println(equiv)
}