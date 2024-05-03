package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "From golang file"

	file, err := os.Create("./a.txt")

	// if err != nil {
	// 	panic(err) //shutsdown the execution of the program and shows u were is the error
	// }
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("Length is", length)
	defer file.Close()
	readFile("./a.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	checkNilError(err)

	fmt.Println("text data in the file is:", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
