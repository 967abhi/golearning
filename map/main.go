package main

import "fmt"

func main() {
	fmt.Println("Hello from the map section")
	students := make(map[string]int)
	students["Abhishek"] = 100
	students["Alice"] = 200
	students["Bob"] = 300
	fmt.Println("Marks of Bob:", students["Bob"])
	fmt.Println("marks of abshishek:", students["Abhishek"])
}
