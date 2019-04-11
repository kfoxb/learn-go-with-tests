package iteration

func Repeat(l string) string {
	var r string
	for i := 0; i < 5; i++ {
		r += l
	}
	return r
}
