package tinyq

import "testing"

func TestLastOfThree(t *testing.T) {
	if got := Last([]int{1, 2, 3}); got != 3 {
		t.Fatalf("Last = %d, want 3", got)
	}
}

func TestSum(t *testing.T) {
	if got := Sum([]int{1, 2, 3}); got != 6 {
		t.Fatalf("Sum = %d, want 6", got)
	}
}
