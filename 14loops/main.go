package main

import (
	"fmt"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// var arr = []int{}

	// for i := 0; i < 3; i++ {
	// 	val, _ := reader.ReadString('\n')
	// 	conVal, _ := strconv.ParseFloat(strings.TrimSpace(val), 64)

	// 	arr = append(arr, int(conVal))

	// }

	// // for _, i := range arr {
	// // 	fmt.Println(i)
	// // }
	// for i := range arr {
	// 	fmt.Println(arr[i])
	// }

	i := 1
	for i < 10 { //same as while loop
		// if i == 5 {
		// 	break
		// }

		// if i == 5 {
		// 	i++
		// 	continue
		// }

		if i == 2 {
			goto lco
		}
		fmt.Println(i)
		i++
	}

lco:
	fmt.Println("Jumping on goto label")

}
