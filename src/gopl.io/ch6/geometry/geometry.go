package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Path []Point

//function
func Distance(p, q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

//method
func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p Path) Distance() float64 {
	sum := 0.0
	for i := range p {
		if i > 0 {
			sum += p[i-1].Distance(p[i])
		}
	}
	return sum
}

func main() {
	p1 := Point{1, 2}
	p2 := Point{3, 4}
	fmt.Println(Distance(p1, p2))
	fmt.Println(p1.Distance(p2))

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1}}
	fmt.Println(perim.Distance())
}
