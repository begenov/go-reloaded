package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
		return
	}
}

func main() {
	file, err := os.ReadFile(os.Args[1])
	check(err)
	strFile := string(file)
	res := strings.Split(strFile, " ")
	// fmt.Print(res)

	for i := 0; i < len(res); i++ {
		if res[i] == "(hex)" {
			res[i-1] = hex(res[i-1])
		}
	}
	fmt.Println(res)
}

func hex(hexa string) string {
	decimal, err := strconv.ParseInt(hexa, 16, 64)
	if err != nil {
		fmt.Println(err)
	}
	return strconv.FormatInt(decimal, 16)
}
