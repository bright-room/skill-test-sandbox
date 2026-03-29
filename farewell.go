package sandbox

// Farewell returns a farewell message for the given name.
func Farewell(name string) string {
	if name == "" {
		return "Goodbye, World!"
	}
	return "Goodbye, " + name + "!"
}
