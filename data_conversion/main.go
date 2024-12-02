package main

import "fmt"

func main() {
	var num int = 42
	fmt.Println("Number is ", num)
	fmt.Printf("Type of num is %T\n", num)
	var data float64 = float64(num)
	fmt.Println("Data is ", data)
	fmt.Printf("Type of data is %T\n", data)
}

// package main

// import "fmt"

// func main() {
// 	var num int = 42
// 	var price float64 = 9.99
// 	var isAvailable bool = true
// 	var name string = "GoLang"

// 	fmt.Printf("Number is and its type is %T\n", num)
// 	fmt.Printf("Price is %.2f and its type is %T\n", price, price)
// 	fmt.Printf("IsAvailable is %v and its type is %T\n", isAvailable, isAvailable)
// 	fmt.Printf("Name is %s and its type is %T\n", name, name)
// }
