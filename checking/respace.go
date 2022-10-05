package checing

func respace(s string) string {
	l := len(s) - 1
	if s[l-1] == ' ' {
		return respace(s[:l])
	}
	return s[:l]
}
