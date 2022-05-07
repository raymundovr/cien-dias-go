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
	fmt.Println("Point C", c)

	var bluePoint geometry.ColoredPoint
	bluePoint.Color = "blue"
	bluePoint.X = 10
	bluePoint.Y = 1.23

	fmt.Println("Punto azul", bluePoint)
}