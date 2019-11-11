package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`
	fields := strings.Fields(text)
	words := make(map[string]int)

	for _, field := range fields {
		words[strings.ToLower(field)]++
	}

	fmt.Println(words)
}
