package checing

func Removeword(s []string) string {
	str := ""
	for i, word := range s {
		if word == "" {
			continue
		} else if s[i] == "(low," || s[i] == "(up," || s[i] == "(cap," {
			s[i] = ""
			s[i+1] = ""
		} else if s[i] != "(up)" && s[i] != "(low)" && s[i] != "(bin)" && s[i] != "(hex)" && s[i] != "(cap)" {
			if i == 0 {
				str += s[i]
			} else {
				str += " " + s[i]
			}
		}
	}
	return str
}

// 3zjad
