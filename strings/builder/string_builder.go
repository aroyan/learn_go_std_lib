package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(strings.Cut("Before", "for"))
	// fmt.Println(strings.Replace("Gunung Bromo", "Bromo", "Semeru", 1))
	// fmt.Println(strings.Replace("Gunung Bromo Bromo Bromo Bromo", "Bromo", "Semeru", 3))
	// fmt.Println(len(strings.ToUpper(strings.Replace(strings.ToLower("Gunung Bromo Bromo Bromo Bromo"), strings.ToLower("BROMO"), "Semeru", -1))))
	// fmt.Println(strings.Replace("Gunung Bromo", "bromo", "Semeru", 122)) // Case sensitive
	// resp, err := http.Get("http://example.com/")

	// if err != nil {
	// 	fmt.Println("Working")
	// 	fmt.Println(resp)
	// } else {
	// 	fmt.Println(err)
	// 	fmt.Println("Its nil")
	// }

	// String builder is building strings from another strings. It's more effective than directly manipulate string
	var sb strings.Builder

	sb.WriteString("This is String frm String Builder\n")
	sb.WriteString("This is String frm String Builder\n")
	sb.WriteString("This is String frm String Builder\n")
	sb.WriteString("This is String frm String Builder\n")

	fmt.Println(sb.String())

	fmt.Println("Capacity", sb.Cap())

	// Set grow capacity
	sb.Grow(1024)

	fmt.Println("Capacity", sb.Cap())

	for i := 0; i <= 10; i++ {
		fmt.Fprint(&sb, "String %d --", i)
	}

	fmt.Println(sb.String())

	fmt.Println(sb.Len())

	sb.Reset()

	fmt.Println(sb.Len())
}
