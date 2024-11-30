package main

import "fmt"

func simpleFunction() {
	fmt.Println("Simple Function")
}
func add(a int, b int) int {
	return a + b
}
func multiply(a, b int) (result int) {
	result = a * b
	return
}
func main() {
	fmt.Println("Hello from the function componenet")
	simpleFunction()
	ans := add(3, 4)
	fmt.Println("add of two number", ans)
	mul := multiply(3, 4)
	fmt.Println("Multiply of two number", mul)

}
