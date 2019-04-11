package iteration

// Repeat creates a string of a character repeated
func Repeat(l string, amount int) string {
	var r string
	for i := 0; i < amount; i++ {
		r += l
	}
	return r
}
