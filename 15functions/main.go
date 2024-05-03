package main

import "fmt"

func main() {
	greet()

	result := add(3, 5)
	fmt.Println("Sum is :", result)

	ans, ansString := proAdd(1, 2, 3, 4, 5)
	fmt.Println("Pro add function sum is :", ans)
	fmt.Println("Pro add function sum is :", ansString)
}

func add(val1 int, val2 int) int {
	return val1 + val2
}

func proAdd(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi from proAdd function"
}

func greet() {
	fmt.Println("Namaste")
}
