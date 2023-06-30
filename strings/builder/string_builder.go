package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	fmt.Println(strings.Cut("Before", "for"))
	fmt.Println(strings.Replace("Gunung Bromo", "Bromo", "Semeru", 1))
	fmt.Println(strings.Replace("Gunung Bromo Bromo Bromo Bromo", "Bromo", "Semeru", 3))
	fmt.Println(len(strings.ToUpper(strings.Replace(strings.ToLower("Gunung Bromo Bromo Bromo Bromo"), strings.ToLower("BROMO"), "Semeru", -1))))
	fmt.Println(strings.Replace("Gunung Bromo", "bromo", "Semeru", 122)) // Case sensitive
	resp, err := http.Get("http://example.com/")

	if err != nil {
		fmt.Println("Working")
		fmt.Println(resp)
	} else {
		fmt.Println(err)
		fmt.Println("Its nil")
	}

}
