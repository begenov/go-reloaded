package main

import (
	"fmt"
	"os"
	"strings"

	checing "go-reloaded/checking"
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
	split_text := strings.Split(text, " ")
	split_text = checing.FirstHint(split_text)
	split_text = checing.Performance(split_text)
	split := checing.Performance(split_text)
	str := checing.Removeword(split)
	str = checing.Punctuation(str)
	str = checing.Sort(str)
	result := []byte(str)
	errResult := os.WriteFile(args[1], result, 0o644)
	checing.Check(errResult)
}

// string to convert []string
