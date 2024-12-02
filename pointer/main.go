// pointer is a variable that stores the address of the value
package main

import "fmt"

func modifyValueByReference(num *int) {
	*num = *num * 20
}

func main() {
	// var num int = 2
	// var ptr *int
	// ptr = &num
	num := "abhishek"
	ptr := &num
	// fmt.Println("Num has values", num)

	fmt.Println("Poninter contains", ptr)
	fmt.Println("Data contains through pointer", *ptr)
	var pointer *int
	if pointer == nil {
		fmt.Println(("Pointer is not assigned"))
	}
	values := 10
	modifyValueByReference(&values)
	fmt.Println("Values contains:", values)
}
