package main

import "fmt"

type Person struct {
	Firstname string
	Lastname  string
	Age       int
}

func main() {
	fmt.Println("I am in structure")
	var abhishek Person
	fmt.Println("Person abhishek:", abhishek)
	abhishek.Firstname = "Abhishek"
	abhishek.Lastname = "Singh"
	abhishek.Age = 32
	fmt.Println("Abhishek firstname is :", abhishek)
	person1 := Person{
		Firstname: "abhinav",
		Lastname:  "singh",
	}
	fmt.Println("person data", person1)

}
