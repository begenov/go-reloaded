package checing

import (
	"fmt"
	"strconv"
	"strings"
)

func Performance(res []string) []string {
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
