package geometry

import (
	"math"
	"fmt"
)

type Point struct {
	X, Y float64	
}

type ColoredPoint struct {
	Point
	Color string
}

func (p Point) Distance(q Point) float64 {
	return math.Sqrt(math.Pow((q.X - p.X), 2) + math.Pow((q.Y - p.Y), 2))
}

func (p Point) ScaleBy(factor float64) Point {
	return Point{ X: p.X * factor, Y: p.Y * factor }
}

func Add(p Point, q Point) Point {
	return Point{ p.X + q.X, p.Y + q.Y }
}

func Substract(p Point, q Point) Point {
	return Point{ p.X - q.X, p.Y - q.Y }
}

func (c ColoredPoint) Distance(q ColoredPoint) float64 {
	return c.Point.Distance(q.Point)
}

func (p Point) String() string {
	return fmt.Sprintf("Point {%.2f, %.2f}", p.X, p.Y)
}

func (c ColoredPoint) String() string {
	return fmt.Sprintf("ColoredPoint %s %s", c.Color, c.Point)
}