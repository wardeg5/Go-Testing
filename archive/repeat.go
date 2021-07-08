package iteration

const repeatcount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatcount; i++ {
		repeated += character
	}
	return repeated
}
