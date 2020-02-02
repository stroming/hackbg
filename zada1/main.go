package main

import (
	"fmt"
	"strings"
)

//Using map I can put a value on "(" = 1 and ")" = -1 and zero value on everything else

var m = map[string]int{
	"(": 1,
	")": -1,
}

func main() {
	d := 0
	res := false
	var b string
	fmt.Scanln(&b)

	a := strings.Split(b, "")

	d = checkCorrOfParens(a)

	res = checkIfTrue(d)

	fmt.Println(res)
}

func checkCorrOfParens(s []string) int {
	d := 0
	for i := 0; i < len(s); i++ {
		d = d + m[s[i]]
		if d < 0 {
			return d
		}
	}

	return d
}

func checkIfTrue(i int) bool {
	res := true
	if i == 0 {
		res = true
	} else {
		res = false
	}
	return res
}
