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
	customers := []Person{
		{"Aroyan", "Bakti", "East Java", 20, "Indonesia"},
		{"John", "Doe", "Amsterdam", 30, "Netherland"},
		{"Taylor", "Swift", "New York", 31, "USA"},
	}

	c, err := json.MarshalIndent(customers, "", "\t")
	fmt.Printf("%s", c)
	fmt.Println(err)

}
