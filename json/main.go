package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:name`
	Age     int    `json age`
	IsAdult bool   `json is_adult`
}

func main() {
	fmt.Println("we are learning json")
	person := Person{
		Name:    "abhishe",
		Age:     3,
		IsAdult: true,
	}
	fmt.Println("find the person", person)
	//Marshaling
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error Marshaling", err)
		return
	}
	fmt.Println("Person data is ", string(jsonData))
	//Unmarshalling
	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error unMarshaling", err)
		return
	}
	fmt.Println("unmarshaling data", personData)

}
