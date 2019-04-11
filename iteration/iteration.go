package iteration

func Repeated(l string) string {
	var r string
	for i := 0; i < 5; i++ {
		r = r + l
	}
	return r
}
