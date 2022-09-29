package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func Check(e error) {
	if e != nil {
		log.Fatal(e)
		return
	}
}

func main() {
	file, err := os.ReadFile(os.Args[1])
	Check(err)
	strFile := string(file)
	res := SplitWithSpaces(strFile)
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
			res[i-1] = ToUpper(res[i-1])
		} else if res[i] == "(low)" {
			res[i-1] = ToLower(res[i-1])
		} else if res[i] == "(cap)" {
			res[i-1] = Capitalize(res[i-1])
		} else if res[i] == "(low," {
			num := res[i+1]
			numn, err := strconv.Atoi(num[:len(num)-1])
			Check(err)
			for j := 0; j < len(res[i-numn:i]); j++ {
				res[(i+j)-numn] = ToLower(res[(i+j)-numn])
			}
		} else if res[i] == "(up," {
			num := res[i+1]
			numn, err := strconv.Atoi(num[:len(num)-1])
			Check(err)
			for j := 0; j < len(res[i-numn:i]); j++ {
				res[(i+j)-numn] = ToUpper(res[(i+j)-numn])
			}

		} else if res[i] == "(cap," {

			num := res[i+1]
			numn, err := strconv.Atoi(num[:len(num)-1])
			Check(err)
			for j := 0; j < len(res[i-numn:i]); j++ {
				res[(i+j)-numn] = Capitalize(res[(i+j)-numn])
			}

		}
	}
	fmt.Println(res)
}

func Capitalize(s string) string {
	str := []byte(ToLower(s))
	ass := true
	for i := 0; i < len(str); i++ {
		if str[i] >= 'a' && str[i] <= 'z' && ass {
			str[i] -= 32
			ass = false
		} else if (str[i] < 'a' || str[i] > 'z') && (str[i] < '0' || str[i] > '9') {
			ass = true
		} else {
			ass = false
		}
	}
	return string(str)
}

func ToLower(s string) string {
	sentence := []rune(s)
	for i := 0; i < len(sentence); i++ {

		currentLetter := sentence[i]

		if currentLetter >= 'A' && currentLetter <= 'Z' {
			sentence[i] += 32
		}
	}
	return string(sentence)
}

func ToUpper(s string) string {
	var str string
	runes := []rune(s)
	for i := range runes {
		if runes[i] >= 'a' && runes[i] <= 'z' {
			str += string(runes[i] - 'a' + 'A')
		}
	}
	return str
}

func SplitWithSpaces(s string) []string {
	var str []string
	var word string
	var booles bool
	for i := range s {
		if s[i] == ' ' && booles {
			str = append(str, word)
			word = ""
			booles = false

		} else if s[i] != ' ' {
			word += string(s[i])
			booles = true
		}
	}
	if len(str) != len(s) {
		str = append(str, word)
	}
	return str
}
