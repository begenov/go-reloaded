package checing

func FirstHint(s []string) []string {
	str := []string{}
	if s[0] == "(cap)" || s[0] == "(up)" || s[0] == "(low)" || s[0] == "(hex)" || s[0] == "(bin)" {
		str = s[1:]
	} else {
		str = s
	}
	return str
}
