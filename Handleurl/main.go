package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning url")
	myurl := "https://jsonplaceholder.typicode.com/todos/"
	fmt.Printf("Type of url %T", myurl)
	paresedUrl, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("Cannot parse url", err)
		return
	}
	fmt.Printf("parese ur types %T\n", paresedUrl)
	fmt.Println("scheme", paresedUrl.Scheme)
	fmt.Println("HOST", paresedUrl.Host)
	fmt.Println("Path", paresedUrl.Path)
	fmt.Println("Rawquery", paresedUrl.RawQuery)

}
