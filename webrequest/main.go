package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning web request service")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting get request", err)
	}
	defer res.Body.Close()
	// fmt.Printf("Types of response ", res)
	//Read the response body
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error found", err)
		return
	}
	fmt.Println("response ", string(data))

}
