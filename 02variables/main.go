package main

import "fmt"

const LoginToken string = "adgadfadf"

func main() {
	var username string = "Prathamesh"
	fmt.Println(username)
	fmt.Printf("Variables is of type : %T \n", username)

	var floatVal float64 = 234.34523567896774866
	fmt.Println(floatVal)
	fmt.Printf("Variable is of type : %T \n", floatVal)

	var intVar int = 0
	fmt.Println(intVar)
	fmt.Printf("Variable is of type : %T \n", intVar)

	//implicit type declaration
	var secondUserName = "prathamesh1"
	fmt.Println(secondUserName)

	//no var style
	//here ':' is walrus operator
	// This syntax is only allowed within any method but outside the method you can't use it
	numberOfUsers := 3000
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)

	var name string = "Prathamesh Theurkar"
	var age int = 20
	fmt.Println("Name : ", name, "\nage : ", age)

	a := 10
	fmt.Println("The value of a : ", a)
}
