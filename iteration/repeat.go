package iteration

func Repeat(str string, repeat int) (repeated string) {
	for i := 0; i < repeat; i++ {
		repeated += str
	}
	return
}
