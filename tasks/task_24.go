package tasks

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p Point) GetX() float64 {
	return p.x
}

func (p Point) GetY() float64 {
	return p.y
}

func GetDistance(p1, p2 Point) float64 {
	return math.Sqrt((p2.x-p1.x)*(p2.x-p1.x) + (p2.y-p1.y)*(p2.y-p1.y))
}

func Task24() {
	p1 := Point{
		x: 1,
		y: 2,
	}

	p2 := Point{
		x: 12,
		y: 14,
	}

	Distance := GetDistance(p1, p2)

	fmt.Println(Distance)
}
