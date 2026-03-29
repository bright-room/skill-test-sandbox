package sandbox

import "testing"

func TestDivide(t *testing.T) {
	got, err := Divide(10, 2)
	if err != nil {
		t.Fatalf("Divide(10, 2) returned unexpected error: %v", err)
	}
	if got != 5 {
		t.Errorf("Divide(10, 2) = %d, want 5", got)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := Divide(10, 0)
	if err == nil {
		t.Error("Divide(10, 0) expected error, got nil")
	}
}

func TestSum(t *testing.T) {
	got := Sum([]int{1, 2, 3})
	if got != 6 {
		t.Errorf("Sum([1,2,3]) = %d, want 6", got)
	}
}
