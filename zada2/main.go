package main

import (
	"fmt"
)

var x int
var y int
var rot int = 3
var x1 int
var y1 int
var x2 int
var y2 int

type location struct {
	i int
	j int
}

func main() {
	// These are the vars that specify how big the Crate is
	fmt.Println("How many apples can the crate store? Please use 'NxN' format for input.:")
	_, err := fmt.Scanf("%dx%d\n", &x, &y)

	// These are the locations of the rotten apples
	fmt.Println("Where's the first rotten apple? Please use '(N,N)'format for input.:")
	_, err1 := fmt.Scanf("(%d,%d)\n", &x1, &y1)
	fmt.Println("Where is the second rotten apple?(If none leave blanc) Please use '(N,N)'format for input,or skip if there's only 1 rotten apple.:")
	_, err2 := fmt.Scanf("(%d,%d)\n", &x2, &y2)

	// How long was I gone
	var timeGone int
	fmt.Println("How many days will you be gone?:")
	_, err3 := fmt.Scanln(&timeGone)
	fmt.Println(err, err1, err2, err3)

	// Creating the matrix that would represent the crate

	a := createMatrix(x, y)

	// Creating crate with the apples in it marked with a 0

	createCrate(a)

	a[x1][y1] = "x"
	a[x2][y2] = "x"

	print(a)

	rotPeriod := timeGone / rot

	rottenApples := []location{}
	// First range over the matrix for the time equal to the rotPeriod
	for p := 0; p < rotPeriod; p++ {
		// Range the actual matrix
		for i := 0; i < len(a); i++ {

			for j := 0; j < len(a[i]); j++ {

				if a[i][j] == "x" {
					// Found the rotten apple
					rottenApples = append(rottenApples, location{i, j})

				}

			}
		}

		for _, apples := range rottenApples {

			ii := apples.i
			jj := apples.j
			if ii-1 >= 0 {
				if jj-1 >= 0 {
					a[ii-1][jj-1] = "x"
				}

				a[ii-1][jj] = "x"

				if jj+1 < len(a) {
					a[ii-1][jj+1] = "x"
				}

			}

			if jj-1 >= 0 {
				a[ii][jj-1] = "x"
			}

			if jj+1 < len(a) {
				a[ii][jj+1] = "x"
			}

			if ii+1 < len(a[0]) {
				if jj-1 >= 0 {
					a[ii+1][jj-1] = "x"
				}
				a[ii+1][jj] = "x"
				if jj+1 < len(a) {
					a[ii+1][jj+1] = "x"
				}
			}

		}

		rottenApples = []location{}
	}
	print(a)

}

func print(m [][]string) {
	fmt.Println()
	for i := 0; i < len(m); i++ {
		fmt.Printf("%v\n", m[i])
	}
}

func createCrate(xi [][]string) {
	for i := 0; i < len(xi); i++ {
		for j := 0; j < len(xi[i]); j++ {
			xi[i][j] = "O"
		}
	}
}

func createMatrix(y int, x int) [][]string {
	a := make([][]string, y)
	for i := range a {
		a[i] = make([]string, x)
	}
	return a
}
