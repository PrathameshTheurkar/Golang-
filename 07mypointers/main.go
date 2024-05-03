package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hello")

	// var ptr *int
	// fmt.Println("Value of ptr is", ptr)
	// OUTPUT : Value of ptr is <nil>

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the value of myNumber : ")
	input, _ := reader.ReadString('\n')

	myNumber, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	}

	// myNumber := 25

	// var ptr *int = &myNumber
	var ptr = &myNumber
	// Refernce means "&"
	fmt.Println("value of actual pointer is", ptr)
	// OUTPUT : value of actual pointer is 0xc00000a0f0
	fmt.Println("value of actual pointer is", *ptr)
	// OUTPUT : value of actual pointer is 25

	*ptr = *ptr * 2
	fmt.Println("New value is", myNumber)

}
