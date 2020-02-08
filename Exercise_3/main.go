package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x   float64
	y   float64
	lvl float64
}

var radius float64
var appropriateError float64
var xi []Point
var p Point
var r int
var data []float64
var points map[string]Point

func main() {
	fmt.Println("Please input Radius:")

	_, err := fmt.Scanln(&r)

	radius = float64(r)

	checkForErrors(err)

	fmt.Println("Please input max Error vaue:")
	_, err1 := fmt.Scanln(&appropriateError)

	checkForErrors(err1)

	points = readFromFile()
	fmt.Println(points)

	brokenSensors(points)

	xi = brokenSensors(points)
	xi = unique(xi)

	fmt.Println("You should check out these Sensors :", xi)

}

func (p Point) Key() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}
func unique(intSlice []Point) []Point {
	keys := make(map[Point]bool)
	list := []Point{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
func brokenSensors(p map[string]Point) []Point {
	for _, v := range p {
		a := v.x
		b := v.y
		d := v.lvl
		for _, c := range p {
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

	return xi
}
func checkForErrors(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func PointFromFile(s string) Point {

	b := strings.Split(s, ",")

	d, err := strconv.ParseFloat(b[0], 64)
	e, err1 := strconv.ParseFloat(b[1], 64)
	r, err2 := strconv.ParseFloat(b[2], 64)
	x := Point{
		x:   d,
		y:   e,
		lvl: r,
	}
	checkForErrors(err)
	checkForErrors(err1)
	checkForErrors(err2)

	return x
}
func readFromFile() map[string]Point {
	var a string

	file, err := os.Open("myfile.txt")
	checkForErrors(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		a = scanner.Text()
		p := PointFromFile(a)
		points[p.Key()] = p

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return points
}
