package main

import (
	"fmt"
	"math"
)

type Point struct {
	x   float64
	y   float64
	lvl float64
}

func main() {
	var xi []Point
	// var ii []Point
	var radius float64
	var appropriateError float64

	radius = 1

	appropriateError = 50

	var p Point

	points := make(map[string]Point)

	data := [][]float64{
		{1, 2, 100},
		{1, 1, 100},
		{5, 5, 20},
		{4, 4, 80},
		{3, 3, 10},
		{3, 5, 0},
		{3, 6, 200},
	}

	for _, v := range data {
		p = Point{
			x:   v[0],
			y:   v[1],
			lvl: v[2],
		}
		points[p.Key()] = p

	}

	for _, v := range points {
		a := v.x
		b := v.y
		d := v.lvl
		for _, c := range points {
			i := c.x
			j := c.y
			f := c.lvl
			if math.Abs(a-i) <= radius || a+radius <= i {
				if math.Abs(b-j) <= radius || (b+radius) >= j {
					if math.Abs(d-f) >= appropriateError {
						xi = append(xi, c)
						// fmt.Println("Y", c)

					}
				}
			}
		}
	}

	fmt.Println(xi)

}

func (p Point) Key() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}
