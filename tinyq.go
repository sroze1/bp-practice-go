// Package tinyq is a tiny queue helper (pipeline end-to-end fixture) with one known bug.
package tinyq

// Last returns the last item of xs.
func Last(xs []int) int {
	return xs[len(xs)-1]
}

// Sum returns the sum of xs.
func Sum(xs []int) int {
	s := 0
	for _, x := range xs {
		s += x
	}
	return s
}
