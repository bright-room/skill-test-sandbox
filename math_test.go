package sandbox

import "testing"

func TestDivide(t *testing.T) {
	got := Divide(10, 2)
	if got != 5 {
		t.Errorf("Divide(10, 2) = %d, want 5", got)
	}
}

func TestSum(t *testing.T) {
	got := Sum([]int{1, 2, 3})
	if got != 6 {
		t.Errorf("Sum([1,2,3]) = %d, want 6", got)
	}
}
