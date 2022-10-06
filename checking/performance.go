package checing

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Performance(s []string) []string {
	for i, word := range s {
		if word == "(hex)" {
			hexd, err := strconv.ParseInt(s[i-1], 16, 64)
			Check(err)
			s[i-1] = strconv.FormatInt(hexd, 10)
			s[i] = ""
		}
		if word == "(bin)" {
			bin, err := strconv.ParseInt(s[i-1], 2, 64)
			Check(err)
			s[i-1] = strconv.FormatInt(bin, 10)
			s[i] = ""
		}
		if word == "(cap)" {
			num := 1
			for j := 1; j <= num; j++ {
				if i-j >= 0 {
					if s[i-j] == "" {
						num++
						continue
					}
					s[i-j] = strings.ToLower(s[i-j])
					s[i-j] = strings.Title(s[i-j])
				}
			}
			s[i] = ""
		}
		if word == "(up)" {
			num := 1
			for j := 1; j <= num; j++ {
				if i-j >= 0 {
					if s[i-j] == "" {
						num++
						continue
					}
					s[i-1] = strings.ToUpper(s[i-1])
				}
			}
			s[i] = ""
		}
		if word == "(low)" {
			num := 1
			for j := 1; j <= num; j++ {
				if i-j >= 0 {
					if s[i-j] == "" {
						num++
						continue
					}
					s[i-1] = strings.ToLower(s[i-1])
				}
			}
			s[i] = ""
		}
		if word == "(cap," {
			if !(s[i+1] == "") {
				runes := []rune(s[i+1])
				str := ""
				str1 := ""
				for k := 0; k < len(runes); k++ {
					if !(runes[k] >= 48 && runes[k] <= 57) && runes[k] != 41 {
						continue
					} else if runes[k] == 41 {
						for l := k + 1; l < len(runes); l++ {
							str1 = str1 + string(runes[l])
						}
					} else {
						str = str + string(runes[k])
					}
				}
				num, err := strconv.Atoi(str)
				Check(err)
				if num <= 0 {
					fmt.Println("Error : number is negative or zero")
				} else {
					if num > i-1 {
						num = i
						if s[i] == ")" {
							num++
						}
					}
					for j := 1; j <= num; j++ {
						if i-j >= 0 {
							if s[i-j] == "" {
								num++
								continue
							}
							s[i-j] = strings.ToLower(s[i-j])
							s[i-j] = strings.Title(s[i-j])
							s[i] = ""
							s[i+1] = str1
						}
					}
				}
			} else {
				fmt.Println("Error: next arguments aren't number")
				os.Exit(0)
			}
		}
		if word == "(low," {
			if !(s[i+1] == "") {
				runes := []rune(s[i+1])
				str := ""
				str1 := ""
				for k := 0; k < len(runes); k++ {
					if !(runes[k] >= 48 && runes[k] <= 57) && runes[k] != 41 {
						continue
					} else if runes[k] == 41 {
						for l := k + 1; l < len(runes); l++ {
							str1 = str1 + string(runes[l])
						}
					} else {
						str = str + string(runes[k])
					}
				}
				num, err := strconv.Atoi(str)
				Check(err)
				if num <= 0 {
					fmt.Println("Error : number is negative or zero")
				} else {
					if num > i-1 {
						num = i
						if s[i] == ")" {
							num++
						}
					}
					for j := 1; j <= num; j++ {
						if i-j >= 0 {
							if s[i-j] == "" {
								num++
								continue
							}
							s[i-j] = strings.ToLower(s[i-j])
							s[i] = ""
							s[i+1] = str1
						}
					}
				}
			} else {
				fmt.Println("Error: next arguments aren't number")
				os.Exit(0)
			}
		}
		if word == "(up," {
			if !(s[i+1] == "") {
				runes := []rune(s[i+1])
				str := ""
				str1 := ""
				for k := 0; k < len(runes); k++ {
					if !(runes[k] >= 48 && runes[k] <= 57) && runes[k] != 41 {
						continue
					} else if runes[k] == 41 {
						for l := k + 1; l < len(runes); l++ {
							str1 = str1 + string(runes[l])
						}
					} else {
						str = str + string(runes[k])
					}
				}
				num, err := strconv.Atoi(str)
				Check(err)
				if num <= 0 {
					fmt.Println("Error : number is negative or zero")
				} else {
					if num > i-1 {
						num = i
						if s[i] == ")" {
							num++
						}
					}
					for j := 1; j <= num; j++ {
						if i-j >= 0 {
							if s[i-j] == "" {
								num++
								continue
							}
							s[i-j] = strings.ToUpper(s[i-j])
							s[i] = ""
							s[i+1] = str1
						}
					}
				}
			} else {
				fmt.Println("Error: next arguments aren't number")
				os.Exit(0)
			}
		}
		if word == "a" {
			fletter := strings.ToLower(s[i+1])
			if fletter[0] == 'a' || fletter[0] == 'e' || fletter[0] == 'i' || fletter[0] == 'o' || fletter[0] == 'u' || fletter[0] == 'h' {
				s[i] = "an"
			}
		}
		if word == "A" {
			fletter := strings.ToLower(s[i+1])
			if fletter[0] == 'a' || fletter[0] == 'e' || fletter[0] == 'i' || fletter[0] == 'o' || fletter[0] == 'u' || fletter[0] == 'h' {
				s[i] = "An"
			}
		}
		if word == "an" {
			fletter := strings.ToLower(s[i+1])
			if (fletter[0] >= 98 && fletter[0] <= 100) || (fletter[0] >= 102 && fletter[0] <= 103) || (fletter[0] >= 106 && fletter[0] <= 110) || (fletter[0] >= 112 && fletter[0] <= 116) || (fletter[0] >= 118 && fletter[0] <= 122) {
				s[i] = "a"
			}
		}
		if word == "An" {
			fletter := strings.ToLower(s[i+1])
			if (fletter[0] >= 98 && fletter[0] <= 100) || (fletter[0] >= 102 && fletter[0] <= 103) || (fletter[0] >= 106 && fletter[0] <= 110) || (fletter[0] >= 112 && fletter[0] <= 116) || (fletter[0] >= 118 && fletter[0] <= 122) {
				s[i] = "A"
			}
		}
	}
	return s
}
