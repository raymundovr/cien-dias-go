package main

import (
	"fmt"
	"vasquezruiz.com/conversions"
)

func main() {
	fmt.Println("**Conversions**")
	c := conversions.Celsius(37)

	fmt.Println(c)

	equiv := c.CalculateEquivalents()

	fmt.Println(equiv)
}