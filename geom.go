package TasksFromTheCourse

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}
type Rectangle struct {
	Width  float64
	Height float64
}
type Triangle struct {
	A float64
	B float64
	C float64
}
type Square struct {
	Side float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (t *Triangle) Area() float64 {
	p := (t.A + t.B + t.C) / 2
	s := math.Sqrt(p * (p - t.A) * (p - t.B) * (p - t.C))
	return s
}

func (t *Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

func (s *Square) Area() float64 {
	return s.Side * s.Side
}

func (s *Square) Perimeter() float64 {
	return 4 * s.Side
}

func PrintShapeInfo(s Shape) {
	fmt.Println("")
	fmt.Printf("Площадь: %.2f\n", s.Area())
	fmt.Printf("Периметр: %.2f\n", s.Perimeter())
}
