package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	lowerTxt := strings.ToLower(text)
	cuttedTxt := strings.Split(lowerTxt, " ")
	var slicedText []string
	for _, word := range cuttedTxt {
		if word != "" {
			slicedText = append(slicedText, word)
		}
	}
	return slicedText
}

func main() {
	fmt.Println("Hello, World!")
}
