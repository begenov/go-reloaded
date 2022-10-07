package checing

func Sort(s string) string {
	str := ""

	removeSpace := false
	for i, letter := range s {
		if i == 0 && letter == '\'' {
			str = str + string(letter)
			removeSpace = true
		} else if letter == 39 && (s[i-1] == ' ') {
			if removeSpace {
				str = str[:len(str)-1]
				str = str + string(letter)
				removeSpace = false
			} else {
				str = str + string(letter)
				removeSpace = true
			}
		} else if i > 1 && s[i-2] == 39 && s[i-1] == ' ' {
			if removeSpace {
				str = str[:len(str)-1]
				str = str + string(letter)
			} else {
				str = str + string(letter)
			}
		} else {
			str = str + string(letter)
			// removeSpace = true
		}
	}
	return str
}
