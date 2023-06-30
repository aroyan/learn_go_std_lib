package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	s := "the quick brown fox jumps over the lazy dog"
	s2 := []string{"one", "two", "three", "four", "five"}
	s3 := "this is a string. With some puctuation, for a demo! Yep."

	f := func(c rune) bool {
		return unicode.IsPunct(c)
	}

	replacer := strings.NewReplacer(".", "|", ",", "|", "!", "|")

	joinResult := strings.Join(s2, "-")
	result2 := strings.Fields(s)
	result3 := strings.FieldsFunc(s3, f)
	result4 := replacer.Replace(s3)

	sub1 := strings.Split(s, " ")
	fmt.Printf("%q\n", sub1)

	sub2 := strings.Split(s, "the")
	fmt.Printf("%q\n", sub2)

	fmt.Printf("%q\n", joinResult)
	fmt.Printf("%q\n", result2)
	fmt.Printf("%q\n", result3)
	fmt.Printf("%q\n", result4)
}
