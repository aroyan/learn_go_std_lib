package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Address     string `json:"address"`
	Age         int    `json:"age"`
	Nationality string `json:"nationality"`
}

func main() {
	cust1 := `{"firstName": "Aroyan", "lastName": "Bakti", "address": "East Java", "age": 20, "nationality": "Indonesia"}`
	cs1 := []byte(cust1)

	var data Person

	json.Unmarshal(cs1, &data)

	fmt.Printf("%+v", data)

}
