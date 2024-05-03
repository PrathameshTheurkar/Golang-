package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You have got 1")
	case 2:
		fmt.Println("You have got 2")
	case 3:
		fmt.Println("You have got 3")
		fallthrough
	case 4:
		fmt.Println("You have got 4")
		fallthrough
	case 5:
		fmt.Println("You have got 5")
	case 6:
		fmt.Println("You have got 6")
	default:
		fmt.Println("what the fuck was that!")
	}
}
