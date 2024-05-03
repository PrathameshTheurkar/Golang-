package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()

	fmt.Println(presentTime)
	// OUTPUT :  2024-03-07 18:11:25.5796408 +0530 IST m=+0.003827501

	fmt.Println(presentTime.Format("02-01-2006"))
	// OUTPUT :	03-07-2024

	fmt.Println(presentTime.Format("02-01-2006 Monday"))
	// OUTPUT :	03-07-2024 Thursday

	fmt.Println(presentTime.Format("02-01-2006 15:04:05"))
	// OUTPUT : 03-07-2024 18:17:02

	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"))
	// OUTPUT :	03-07-2024 18:17:40 Thursday

	createdDate := time.Date(2023, time.June, 3, 18, 23, 23, 0, time.UTC)
	fmt.Println("CreateDate is", createdDate)
	// OUTPUT : CreateDate is 2023-06-03 18:23:23 +0000 UTC

	fmt.Println(createdDate.Format("02-01-2006 Monday"))
	// OUTPUT : 03-06-2023

	// go env

	// GOOS="linux" go build
	// this will build this time application for linux OS

	// GOOS="darwin" go build
	// this will build this time application for MAC OS

	// GOOS="windows" go build
	// this will build this time application for windows OS
}
