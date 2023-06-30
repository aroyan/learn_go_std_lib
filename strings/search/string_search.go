package main

import (
	"fmt"
	"strings"
)

func main() {

	fname := "golang.txt"
	fname2 := "temp_gopher.png"
	fname3 := "Tempgopher.png"
	vowels := "aiueoAIUEO"
	s := "The quick brown fox jumps over the lazy dog"

	// Check given string if it contains the given substring
	fmt.Println(strings.Contains(s, "fox")) // Case sensitive
	fmt.Println(strings.Contains(s, "cat"))

	// Check index of given substring in the given string
	fmt.Println(strings.Index(s, "lazy"))
	fmt.Println(strings.Index(s, "Lazy")) // Case sensitive

	// Check if the given strings is match any of given chars
	fmt.Println(strings.IndexAny("bgbhbfh", vowels)) // Return -1 if doesn't match any of the vowels
	fmt.Println(strings.IndexAny("Golang", vowels))  // Return 1 if it match

	// Check strings if it contains suffix or prefix
	fmt.Println(strings.HasSuffix(fname, "txt"))
	fmt.Println(strings.HasPrefix(fname2, "temp"))
	fmt.Println(strings.HasPrefix(fname3, "temp")) // Case sensitive

	// Count given substring appearances in the given string
	fmt.Println(strings.Count(s, "the")) // Case sensitive
	fmt.Println(strings.Count(s, "he"))
}
