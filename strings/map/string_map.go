package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "The quick brown fox jumps over the lazy dog"
	shift := 2

	// Create mapping functions
	transform := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			value := int('A') + (int(r) - int('A') + shift)
			if value > 91 {
				value -= 26
			} else if value < 65 {
				value += 26
			}
			return rune(value)
		case r >= 'a' && r <= 'z':
			value := int('a') + (int(r) - int('a') + shift)
			if value > 122 {
				value -= 26
			} else if value < 97 {
				value += 26
			}
			return rune(value)
		}
		return r
	}

	encode := strings.Map(transform, s)
	fmt.Println(encode)

	shift = -shift
	decode := strings.Map(transform, encode)
	fmt.Println(decode)
}
