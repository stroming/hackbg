package main

import "fmt"

type Point struct {
	x   int
	y   int
	lvl int
}

func main() {
	var xi []Point
	// var ii []Point
	radius := 1
	appropriateError := 50
	var p Point
	points := make(map[string]Point)

	data := [][]int{
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
			if a-radius >= i || a+radius <= i {
				if b-radius >= j || b+radius <= j {
					if d != f {
						if d-f >= appropriateError || f-d >= appropriateError {
							xi = append(xi, c)
							// fmt.Println("Y", c)
						}
					}
				}
			}
		}
	}

	// for _, v := range xi {

	// }
	fmt.Println(xi)
	// k := fmt.Sprint("%d:%d", p.x, p.y+radius)
	// neighbourPoint, ok := points[k]
	// if ok {

	// 	// if ok == true , there is point found
	// 	// see if the two points have huuuge difference between level.
	// }

}

func (p Point) Key() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

// func (p Point) readFromFile(filename string) deck{
// 	bs, err := ioutil.ReadFile(filename)
// 	if err!=nil {
// 		fmt.Println("Error :", err)
// 		os.Exit(1)
// 	}

// return Point(s)

// }

// package main

// import "fmt"

// var cases = []struct {
// 	i             int
// 	j             int
// 	polutionLevel int
// }{
// 	{
// 		i:             3,
// 		j:             4,
// 		polutionLevel: 100,
// 	}, {
// 		i:             3,
// 		j:             3,
// 		polutionLevel: 100,
// 	}, {
// 		i:             4,
// 		j:             3,
// 		polutionLevel: 20,
// 	}, {
// 		i:             10,
// 		j:             9,
// 		polutionLevel: 200,
// 	}, {
// 		i:             9,
// 		j:             9,
// 		polutionLevel: 180,
// 	}, {
// 		i:             7,
// 		j:             6,
// 		polutionLevel: 50,
// 	}, {
// 		i:             7,
// 		j:             5,
// 		polutionLevel: 300,
// 	}, {
// 		i:             1,
// 		j:             1,
// 		polutionLevel: 300,
// 	},
// }

// func main() {
// 	// radius := 1

// 	a := createMatrix()

// 	locationOfSensors(a)

// 	for i := range a {
// 		for j := range a[i] {
// 			if a[i][j] != 0 {

// 			}
// 		}

// 	}

// 	// for _, v := range cases {
// 	// 	if v.i-radius >= 0 {
// 	// 		for i := v.i - radius; i < len(a[v.i-radius:v.i+radius]); i++ {
// 	// 			fmt.Println(a[i])
// 	// 		}
// 	// 	}
// 	// 	if v.i-radius < 0 {

// 	// 		for i := 0; i < len(a[0:v.i+radius+1]); i++ {
// 	// 			fmt.Println(a[i])
// 	// 		}
// 	// 	}
// 	// }

// 	// print(a)

// }

// func createMatrix() [][]int {
// 	a := make([][]int, 12)
// 	for i := range a {
// 		a[i] = make([]int, 12)
// 	}
// 	return a
// }

// func print(s [][]int) {
// 	for i := range s {
// 		fmt.Printf("%v\n", s[i])
// 	}
// }

// func locationOfSensors(a [][]int) {
// 	for _, v := range cases {
// 		a[v.i][v.j] = v.polutionLevel
// 	}
// }
