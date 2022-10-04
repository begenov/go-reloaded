package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 || !(len(args[0]) >= 4 && args[0][len(args[0])-4:] == ".txt") || !(len(args[1]) >= 4 && args[1][len(args[1])-4:] == ".txt") {
		fmt.Println("Incorrect Input")
		return
	}
	file, err := os.ReadFile(args[0])
	Check(err)
	text := string(file)
	// text = sort(text)

	split_text := strings.Fields(text)
	// fmt.Println(mas)
	/*for _, v := range text {
		fmt.Println(string(v))
	}*/
	//  split_text := strings.Split(text, " ")
	split_text = firstHint(split_text)
	split_text = performance(split_text)
	str := removeword(split_text)
	str = Punctuation(str)
	// str = SortApostr(str)
	// fmt.Println(str)
	result := []byte(str)
	errResult := os.WriteFile(args[1], result, 0644)
	Check(errResult)
}

// func space(str string) string {
// 	s := ""
// 	for i := range str {
// 		if str[i] == ' ' {
// 			s += "" + str
// 		} else {
// 			s = " " + str
// 		}
// 	}
// 	return s
// }

func performance(res []string) []string {
	for i := range res {
		if res[i] == "(hex)" {
			hex, err := strconv.ParseInt(res[i-1], 16, 64)
			Check(err)
			res[i-1] = strconv.FormatInt(hex, 10)
		} else if res[i] == "(bin)" {
			if res[i-1] == "" {
				fmt.Println("Incorect input")
				break
			}
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
			if numn <= 0 {
				fmt.Println("Error : number is negative or zero")
				continue
			} else {
				if numn > i-1 {
					numn = i
				}
				for j := 1; j <= numn; j++ {
					if i-j > 0 {
						if res[i-j] == "" {
							numn++
							continue
						}
						if !(res[i-j] >= string(33) && res[i-j] <= string(64)) || (res[i-j] >= string(91) && res[i-j] <= string(96)) || (res[i-j] >= string(123) && res[i-j] <= string(126)) {
							res[i-j] = strings.ToLower(res[i-j])
						} else {
							numn++
						}
					}
				}
			}
		} else if res[i] == "(up," {
			num := res[i+1]
			numn, err := strconv.Atoi(num[:len(num)-1])
			Check(err)
			if numn <= 0 {
				fmt.Println("Error : number is negative or zero")
			} else {

				if numn > i-1 {
					numn = i
				}
				for j := 1; j <= numn; j++ {
					if i-j > 0 {
						if res[i-j] == "" {
							numn++
							continue
						}
						if !(res[i-j] >= string(33) && res[i-j] <= string(64)) || (res[i-j] >= string(91) && res[i-j] <= string(96)) || (res[i-j] >= string(123) && res[i-j] <= string(126)) {
							res[i-j] = strings.ToUpper(res[i-j])
						} else {
							numn++
						}
					}
				}
			}

		} else if res[i] == "(cap," {
			num := res[i+1]
			numn, err := strconv.Atoi(num[:len(num)-1])
			Check(err)
			if numn <= 0 {
				fmt.Println("Error : number is negative or zero")
			} else {
				if numn > i-1 {
					numn = i
				}
				for j := 1; j <= numn; j++ {

					if res[i-j] == "" {
						numn++
						continue
					}
					if !(res[i-j] >= string(33) && res[i-j] <= string(64)) || (res[i-j] >= string(91) && res[i-j] <= string(96)) || (res[i-j] >= string(123) && res[i-j] <= string(126)) {
						res[i-j] = strings.ToLower(res[i-j])
						res[i-j] = strings.Title(res[i-j])
					} else {
						numn++
					}

				}
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
		} else if res[i] == "an" {
			low := strings.ToLower(res[i+1])
			if low[0] == 'a' || low[0] == 101 || low[0] == 'u' || low[0] == 'e' || low[0] == 'o' || low[0] == 'h' {
				res[i] = "a"
			}
		} else if res[i] == "An" {
			low := strings.ToLower(res[i+1])
			if low[0] == 'a' || low[0] == 101 || low[0] == 'u' || low[0] == 'e' || low[0] == 'o' || low[0] == 'h' {
				res[i] = "A"
			}
		}
	}
	return res
}

func Check(err error) {
	if err != nil {

		fmt.Println("ERROR:", err)
		os.Exit(0)
	}
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

func sort(s string) string {
	str := ""

	removeSpace := false
	for i, letter := range s {
		if letter == 39 && (s[i-1] == ' ') {
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
		}
	}
	return str
}

func firstHint(s []string) []string {
	str := []string{}
	if s[0] == "(cap)" || s[0] == "(up)" || s[0] == "(low)" || s[0] == "(hex)" || s[0] == "(bin)" {
		str = s[1:]
	} else {
		str = s
	}
	return str
}
