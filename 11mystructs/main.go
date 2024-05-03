package main

import "fmt"

func main() {
	// no inheritance in golang ; no super or parent or child concept in golang

	prathamesh := User{"Prathamesh", "prathameshtheurkar037@gmail.com", true, 20}
	fmt.Println(prathamesh)
	fmt.Printf("%T\n", prathamesh)
	fmt.Printf("prathamesh details are %+v\n", prathamesh)
	fmt.Printf("Name is %v\n", prathamesh.Name)

	// user := User{}
	// user.Name = "xyc"
	// fmt.Printf("%+v", user)

	var todos = []Todo{}
	todo := Todo{"prathamesh", "theurkar"}
	todos = append(todos, todo)
	fmt.Println(todos)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

type Todo struct {
	Title       string
	Description string
}
