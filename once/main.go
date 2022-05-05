package main

import (
	"fmt"
	"vasquezruiz.com/geometry"
)

func main() {
	p := geometry.Point{X: 2.43, Y: 3.1}
	q := geometry.Point{X: 1.46, Y: 5.41}

	fmt.Printf("Distance from P %s to Q %s is %.2f\n", p, q, p.Distance(q))
	
	color := "red"
	c := geometry.ColoredPoint{ geometry.Point{1, 1}, color }

	fmt.Printf("Point C %s", c)
}