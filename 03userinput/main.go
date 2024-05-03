package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to userInput"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	// comma ok  or  err ok syntax
	input, _ := reader.ReadString('\n')
	// input, err := reader.ReadString('\n')
	// _, err := reader.ReadString('\n')
	fmt.Print(input)
	fmt.Printf("Type of input : %T\n", input)

	// reader1 := bufio.NewReader(os.Stdin)
	username, _ := reader.ReadString('\n')
	fmt.Println("UserName is", username)

}
