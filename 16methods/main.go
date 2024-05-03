package main

import "fmt"

func main() {
	// no inheritance in golang ; no super or parent or child concept in golang

	prathamesh := User{"Prathamesh", "prathameshtheurkar037@gmail.com", true, 20}
	fmt.Println(prathamesh)
	fmt.Printf("%T\n", prathamesh)
	fmt.Printf("prathamesh details are %+v\n", prathamesh)
	fmt.Printf("Name is %v\n", prathamesh.Name)
	prathamesh.GetStatus()
	prathamesh.NewMail()
	fmt.Printf("prathamesh details are %+v\n", prathamesh)

	fmt.Println(prathamesh.PrintInfo(2))

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u *User) GetStatus() {
	fmt.Println("Is user active :", u.Status)
}

func (u *User) NewMail() {
	u.Email = "prathamesh.theurar21@pccoepune.org"
	fmt.Println("Email of User is", u.Email)
}

func (u *User) PrintInfo(a int) int {
	fmt.Println("Prathamesh Theurkar")
	return a * 2
}
