package hamming

import "fmt"

// Distance calculates the hammings distance between 2 dna strings.If the given dnas are of different length return an error
func Distance(a, b string) (int, error) {
	res := 0
	if len(a) != len(b) {
		return res, fmt.Errorf("len(a): %v is not equal to len(b): %v", len(a), len(b))
	}
	for i := 0; i < len(a); i++ {
		if b[i] != a[i] {
			res++
		}
	}
	return res, nil
}
