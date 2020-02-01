package main

import "fmt"

const (
	//input za
	x int = 5
	y int = 5
	//n na broi
	//y = kolona x = red
	x1 int = 3
	y1 int = 3

	// x2 int = 3
	// y2 int = 5

	// x3 int = 9
	// y3 int = 1
)

type matrix [x][y]string

var a matrix

func main() {
	// timeGone := 2

	//tova e za kasetkata
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			a[i][j] = "G"
		}
	}
	//tova sa gnilite qbulki koordinati
	a[y1][x1] = "x"

	// a[y2][x2] = "x"
	// a[y3][x3] = "x"
	//Tuka susednite na a[y1][x2] se prevrushtata v x=sove i nagore i na dolo
	// if timeGone > 3 {
	if y1 > 1 {
		for i := y1 - 1; i < len(a[y1-1:y1+2]); i++ {
			if x1 > 1 {
				for j := x1 - 1; j < len(a[x1-1:x1+2]); j++ {
					a[i-1 : i+2][j-2 : j+2] = "x"
				}
			} else if x1 == 1 {

				for j := x1; j < len(a[x1:x1+2]); j++ {
					a[i][j] = "x"
				}
			}

		}

	} else if y1 == 1 {
		for i := y1; i < len(a[y1:y1+2]); i++ {
			if x1 > 1 {
				for j := x1 - 1; j < len(a[x1-1:x1+2]); j++ {
					a[i][j] = "x"
				}
			} else if x1 == 1 {

				for j := x1; j < len(a[x1:x1+2]); j++ {
					a[i][j] = "x"
				}
			}
		}

	}

	// } else if timeGone > 6 {
	// 	for i := y1 - 1; i < len(a[y1-3:y1+4]); i++ {
	// 		for j := x1 - 1; j < len(a[x1-3:x1+4]); j++ {
	// 			a[i][j] = "x"

	// 		}
	// 	}

	// }

	// for i := 0; i < len(a); i++ {
	// 	for j := 0; j < len(a[i]); j++ {
	// 		if a[i][j] == "x" {
	// 			fmt.Println("yes bro fck yes")
	// 			a[i-1][j-1] = "x"

	// 		}
	// 	}
	// }

	//TOVA E DA VIDISH KAK 	izgLEJDA KATO MATRICA FNIMAAI BRATO !!!!!!
	for i := 0; i < len(a); i++ {
		fmt.Printf("%v\n", a[i])
	}

}
