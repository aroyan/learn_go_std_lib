package main

import (
	"fmt"
	"unicode"
)

func main() {
	const s = "The Quick brown fox 'Jump' Over the lazy *Dogs*"

	punctuationCount := 0
	lowercaseCount := 0
	uppercaseCount := 0
	spaceCount := 0
	hexDigitCount := 0

	for _, ch := range s {
		if unicode.IsPunct(ch) {
			punctuationCount++
		}

		if unicode.IsLower(ch) {
			lowercaseCount++
		}

		if unicode.IsUpper(ch) {
			uppercaseCount++
		}

		if unicode.IsSpace(ch) {
			spaceCount++
		}

		if unicode.Is(unicode.Hex_Digit, ch) {
			hexDigitCount++
		}
	}

	fmt.Println("Punctuation count:", punctuationCount)
	fmt.Println("Lowercase count:", lowercaseCount)
	fmt.Println("Uppercase count:", uppercaseCount)
	fmt.Println("Space count:", spaceCount)
	fmt.Println("Heximal digit count:", hexDigitCount)
}
