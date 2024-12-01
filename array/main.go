package main

import "fmt"

func main() {
	fmt.Println("we are learning array in golang")
	var name [5]string
	name[0] = "Abhishek"

	name[2] = "abhinav"
	name[3] = "raj"

	fmt.Println("Names of person is :", name)
	fmt.Println("Lengrth of number array is :", len(name))
	fmt.Println("ser the value :", name[2])
	fmt.Printf("name array is %q\n", name)
}
