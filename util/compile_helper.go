package util

// Use masks unused variables when compiling go program.
func Use(vals... interface{}) {
	for _, val := range vals {
		_ = val
	}
}
