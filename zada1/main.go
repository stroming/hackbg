package main

import (
	"fmt"
	"strings"
)

var res bool
var d int

//Using map I can put a value on "(" = 1 and ")" = -1 and zero value on everything else

var m = map[string]int{
	"(": 1,
	")": -1,
}

func main() {
	b := "(()()())()"

	a := strings.Split(b, "")

	d = checkCorrOfParens(a)

	fmt.Printf("This is main %d\n", d)

	checkIfTrue(d)

	fmt.Println(res)
}

func checkCorrOfParens(s []string) int {
	for i := 0; i < len(s); i++ {
		fmt.Printf("This is 1st %d\n", d)
		d = d + m[s[i]]
		fmt.Printf("This is 2nd %d\n", d)
		if i == 0 && d != 1 {
			break
		} else if i == (len(s)-1) && m[s[i]] == 1 {
			d = 10
		}
	}
	fmt.Printf("This is at the end %d\n", d)
	return d
}

func checkIfTrue(i int) bool {
	if d == 0 {
		res = true
	} else {
		res = false
	}
	return res
}
