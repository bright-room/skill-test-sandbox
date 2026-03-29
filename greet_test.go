package sandbox

import "testing"

func TestGreet(t *testing.T) {
	tests := []struct {
		name string
		input string
		want string
	}{
		{"with name", "Go", "Hello, Go!"},
		{"empty name", "", "Hello, World!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greet(tt.input); got != tt.want {
				t.Errorf("Greet(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
