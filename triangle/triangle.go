// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	Equ Kind = iota
	Iso Kind = iota
	Sca Kind = iota
	NaT Kind = iota
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {

	// a > 0 b>0 c>0
	// a+b >= c c+b >=a c+a >= b
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}
	if math.IsInf(a, -1) || math.IsInf(b, -1) || math.IsInf(c, -1) {
		return NaT
	}
	if math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) {
		return NaT
	}
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}
	if a+b < c || c+b < a || c+a < b {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a != b && b != c && c != a {
		return Sca
	}
	if (a == b && b != c) || (b == c && c != a) || (a == c && c != b) {
		return Iso
	}

	return NaT
}
