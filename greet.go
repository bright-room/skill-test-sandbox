package sandbox

// Greet returns a greeting message for the given name.
func Greet(name string) string {
	if name == "" {
		return "Hello, World!"
	}
	return "Hello, " + name + "!"
}
