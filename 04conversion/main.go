package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the value of val : ")
	val, _ := reader.ReadString('\n')
	fmt.Println("value is", val)

	conVal, err := strconv.ParseFloat(strings.TrimSpace(val), 64)

	if err != nil {
		fmt.Println(err)
		// panic(err)
		// This basically on error shutdowns the program
	} else {
		fmt.Println("Increasesd val is :", conVal+1)
	}

}

func checkNilError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
