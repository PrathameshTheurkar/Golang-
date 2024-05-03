package main

import (
	"io/ioutil"
	"os"
)

func main() {
	// content, err := ioutil.ReadFile("./a.txt")
	// checkNilError(err)
	// fmt.Println("content is:", content)
	// // ioutil.WriteFile("./b.txt", content, 0644)

	// readFile, err := os.Open("./a.txt")
	// checkNilError(err)
	// defer readFile.Close()

	// writeFile, err := os.Create("output.txt")
	// checkNilError(err)
	// defer writeFile.Close()

	// length, err := io.Copy(writeFile, readFile)
	// fmt.Println("Length:", length)

	// appending the one file content to the existing file

	f, err := os.OpenFile("./output.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	checkNilError(err)

	content, err := ioutil.ReadFile("./a.txt")
	if _, err := f.WriteString(string(content)); err != nil {
		panic(err)
	}

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
