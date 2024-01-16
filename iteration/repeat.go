package iteration

const repeatCount = 5

func Repeat(s string) string {
	c := ""
	for i := 0; i < repeatCount; i++ {
		c += s
	}
	return c
}
