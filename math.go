package sandbox

// Divide returns a / b.
func Divide(a, b int) int {
	return a / b
}

// Sum returns the sum of numbers.
func Sum(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum = sum + numbers[i]
	}
	return sum
}
