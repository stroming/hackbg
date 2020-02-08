package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
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

func main() {
	fmt.Println("Please input Radius:")

	_, err := fmt.Scanln(&r)

	radius = float64(r)

	checkForErrors(err)

	fmt.Println("Please input max Error vaue:")
	_, err1 := fmt.Scanln(&appropriateError)

	checkForErrors(err1)

	points := make(map[string]Point)
	x := readFromFile()
	data = 

	// data = [][]float64{
	// 	{1, 2, 100},
	// 	{1, 1, 100},
	// 	{5, 5, 20},
	// 	{4, 4, 80},
	// 	{3, 3, 10},
	// 	{3, 5, 0},
	// 	{3, 6, 200},
	// }

	for _, v := range data {
		p = Point{
			x:   v[0],
			y:   v[1],
			lvl: v[2],
		}
		points[p.Key()] = p

	}

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
func brokenSensors(points map[string]Point) []Point {
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

	return xi
}
func checkForErrors(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func DataFromFile() []float64 {
	var x []float64
	a := readFromFile()
	for _, v := range a {
		b := strings.Split(v, ",")
		for _, g := range b {
			d, err := strconv.ParseFloat(g, 64)
			if err != nil {
				fmt.Printf("Error : %v\n", err)
			}
			x = append(x, d)

		}

	}
	return x
}
func readFromFile() []string {
	var a string
	var b []string
	file, err := os.Open("myfile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		a = scanner.Text()
		b = append(b, a)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return b
}
