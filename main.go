package main

import (
	"fmt"
	checing "go-reloaded/checking"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 || !(len(args[0]) >= 4 && args[0][len(args[0])-4:] == ".txt") || !(len(args[1]) >= 4 && args[1][len(args[1])-4:] == ".txt") {
		fmt.Println("Incorrect Input")
		return
	}
	file, err := os.ReadFile(args[0])
	checing.Check(err)
	text := string(file)
	split_text := strings.Fields(text)
	split_text = checing.FirstHint(split_text)
	split_text = checing.Performance(split_text)
	str := checing.Removeword(split_text)
	str = checing.Punctuation(str)
	result := []byte(str)
	errResult := os.WriteFile(args[1], result, 0644)
	checing.Check(errResult)
}
