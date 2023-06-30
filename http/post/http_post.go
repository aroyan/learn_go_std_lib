package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	base_url := "https://jsonplaceholder.typicode.com/todos"

	req_body := `{"userId": 10, "title": "Learning Go"}`

	reqBody := bytes.NewBuffer([]byte(req_body))

	res, err := http.Post(base_url, "application/json", reqBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	fmt.Println(sb)
}
