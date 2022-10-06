package checing

func Punctuation(s string) string {
	str := ""
	for i, l := range s {
		if i == len(s)-1 {
			if l == '.' || l == ',' || l == '!' || l == '?' || l == ':' || l == ';' {
				if s[i-1] == ' ' || s[i-1] == '\n' {
					str = respace(str)
					str = str + string(l)
				} else {
					str = str + string(l)
				}
			} else {
				str = str + string(l)
			}
		} else if l == '.' || l == ',' || l == '!' || l == '?' || l == ':' || l == ';' {
			if s[i-1] == ' ' || s[i-1] == '\n' {
				str = respace(str)
				str = str + string(l)
			} else {
				str = str + string(l)
			}
			if s[i+1] != ' ' && s[i+1] != '.' && s[i+1] != ',' && s[i+1] != '!' && s[i+1] != '?' && s[i+1] != ':' && s[i+1] != ';' {
				str = str + " "
			}
		} else {
			str = str + string(l)
		}
	}
	return str
}
