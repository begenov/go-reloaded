package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile(os.Args[1])
	Check(err)
	str := string(file)
	word := strings.Split(str, " ")
	word = Running(word)
	// fmt.Println(word)
	strword := removeword(word)
	strword = Punctuation(strword)
	// fmt.Println(strword)
	res := []byte(strword)
	out := os.Args[2]
	errRes := os.WriteFile(out, res, 0644)
	Check(errRes)
}

func Running(res []string) []string {
	for i := range res {
		if res[i] == "(hex)" {
			hex, err := strconv.ParseInt(res[i-1], 16, 64)
			Check(err)
			res[i-1] = strconv.FormatInt(hex, 10)
		} else if res[i] == "(bin)" {
			bin, err := strconv.ParseInt(res[i-1], 2, 64)
			Check(err)
			res[i-1] = strconv.FormatInt(bin, 10)
		} else if res[i] == "(up)" {
			res[i-1] = strings.ToUpper(res[i-1])
		} else if res[i] == "(low)" {
			res[i-1] = strings.ToLower(res[i-1])
		} else if res[i] == "(cap)" {
			res[i-1] = strings.Title(res[i-1])
		} else if res[i] == "(low," {
			num := res[i+1]
			numn, err := strconv.Atoi(num[:len(num)-1])
			Check(err)
			for j := 0; j < len(res[i-numn:i]); j++ {
				res[(i+j)-numn] = strings.ToLower(res[(i+j)-numn])
			}
		} else if res[i] == "(up," {
			num := res[i+1]
			numn, err := strconv.Atoi(num[:len(num)-1])
			Check(err)
			for j := 0; j < len(res[i-numn:i]); j++ {
				res[(i+j)-numn] = strings.ToUpper(res[(i+j)-numn])
			}

		} else if res[i] == "(cap," {

			num := res[i+1]
			numn, err := strconv.Atoi(num[:len(num)-1])
			Check(err)
			for j := 0; j < len(res[i-numn:i]); j++ {
				res[(i+j)-numn] = strings.Title(res[(i+j)-numn])
			}

		} else if res[i] == "a" {

			low := strings.ToLower(res[i+1])
			if low[0] == 'h' || low[0] == 'a' || low[0] == 'e' || low[0] == 'i' || low[0] == 'o' || low[0] == 'u' {
				res[i] = "an"
			}

		} else if res[i] == "A" {
			low := strings.ToLower(res[i+1])
			if low[0] == 'h' || low[0] == 'a' || low[0] == 'e' || low[0] == 'i' || low[0] == 'o' || low[0] == 'u' {
				res[i] = "An"
			}
		}
	}
	return res
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
		return

	}
}

func removeword(s []string) string {
	str := ""

	for i := range s {
		if s[i] == "(low," || s[i] == "(up," || s[i] == "(cap," {
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

func Punctuation(s string) string {
	str := ""

	for i, l := range s {
		if i == len(s)-1 {
			if l == '.' || l == ',' || l == '!' || l == '?' || l == ':' || l == ';' {
				if s[i-1] == ' ' {
					str = ReSpace(str)
					str = str + string(l)
				} else {
					str = str + string(l)
				}
			} else {
				str = str + string(l)
			}
		} else if l == '.' || l == ',' || l == '!' || l == '?' || l == ':' || l == ';' {
			if s[i-1] == ' ' {
				str = ReSpace(str)
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

func ReSpace(s string) string {
	l := len(s) - 1
	if s[l-1] == ' ' {
		return ReSpace(s[:l])
	}
	return s[:l]
}
