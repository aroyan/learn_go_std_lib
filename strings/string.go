package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "the quick brown fox jumps over the lazy dog"

	fmt.Println(strings.Compare("hello", "Hello"))

	fmt.Println(strings.Compare("hello", "xxx"))

	x := strings.Count("Loem ipsum dolor sit amet", "Amet")

	fmt.Print(x)

	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.Title(s)) // Deprecated
}
