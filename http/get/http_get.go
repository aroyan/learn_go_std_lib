package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// base_url := "https://httpbin.org"
	base_url := "https://jsonplaceholder.typicode.com/todos"

	res, err := http.Get(base_url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	stringBody := string(body)

	fmt.Println(res.Status)
	fmt.Println(res.StatusCode)
	fmt.Printf("%s", stringBody)
}
