package main

import (
	"fmt"
	"strconv"
)

func main() {

	sampleInt := 20
	sampleString := "250"

	fmt.Println("1")

	sInt := strconv.Itoa(sampleInt)

	fmt.Printf("%T %v \n", sInt, sInt)

	sStr, err := strconv.Atoi(sampleString)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%T %v\n", sStr, sStr)

	// Other parse func
	b, _ := strconv.ParseBool("true")
	fmt.Println(b)

	f, _ := strconv.ParseFloat("2.0123", 64)
	fmt.Println(f)

	// Other format func
	s := strconv.FormatBool(true)

	fmt.Printf("%T, %v", s, s)
}
