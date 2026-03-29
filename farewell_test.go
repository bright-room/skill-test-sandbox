package sandbox

import "testing"

func TestFarewell(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"with name", "Go", "Goodbye, Go!"},
		{"empty name", "", "Goodbye, World!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Farewell(tt.input); got != tt.want {
				t.Errorf("Farewell(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
