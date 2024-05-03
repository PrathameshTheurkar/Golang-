package main

import "fmt"

func main() {
	var arr [4]int

	arr[0] = 1
	arr[1] = 2
	arr[3] = 3

	fmt.Println(arr)
	fmt.Println("Length of the arr is", len(arr))

	var fruitList = [3]string{"Apple", "Banana", "Mango"}
	fmt.Println(fruitList)
}
