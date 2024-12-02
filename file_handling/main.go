package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// file, err := os.Create("example.txt")
	// if err != nil {
	// 	fmt.Println("Error while creating file", err)
	// 	return
	// }
	// defer file.Close()
	// content := "hello world by abhishek"
	// byte, errors := io.WriteString(file, content+"\n")
	// fmt.Println("Bytes of the code", byte)
	// if errors != nil {
	// 	fmt.Println("Error while writing file:", errors)
	// }
	// fmt.Println("successfull created file")

	// file, err := os.Open("example.txt")
	// if err != nil {
	// 	fmt.Println("Error while opening file", err)
	// 	return
	// }
	// buffer := make([]byte, 1024)
	// //Read the file content from the  buffer
	// for {
	// 	n, err := file.Read(buffer)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println("Error while reading file", err)
	// 		return
	// 	}
	// 	//process the read content
	// 	fmt.Println(string(buffer[:n]))
	// }
	content, err := ioutil.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error while reading the file", err)
		return
	}
	fmt.Println(string(content))

}
