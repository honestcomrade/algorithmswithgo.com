package module01

// Reverse will return the provided word in reverse
// order. Eg:

// Reverse("cat") => "tac"
// Reverse("alphabet") => "tebahpla"
func Reverse(word string) string {
	s := make([]rune, 0)
	for _, c := range word {
		s = prependStr(s, c)
	}

	return string(s)
}

func prependStr(x []rune, y rune) []rune {
	x = append(x, 0)
	copy(x[1:], x)
	x[0] = y
	return x
}
