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
	fmt.Print("Enter the value of loginCount: ")
	loginCount1, _ := reader.ReadString('\n')

	loginCount, _ := strconv.ParseFloat(strings.TrimSpace(loginCount1), 64)
	var result string
	if loginCount < 10 {
		result = "Regulate user"
	} else if loginCount > 10 {
		result = "watch out"
	} else {
		result = "somethig else"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is greater than or equal to 10")
	}

}
