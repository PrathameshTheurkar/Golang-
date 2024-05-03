package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Banana", "Mango"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)
	fmt.Println(fruitList)
	fmt.Println("Length of fruitList is", len(fruitList))

	fruitList = append(fruitList, "Peaches", "Berries")
	fmt.Println(fruitList)

	// fruitList = append(fruitList[0:])
	// OUTPUT : [Apple Banana Mango Peaches Berries]

	// fruitList = append(fruitList[1:])
	// OUTPUT :[Banana Mango Peaches Berries]

	fruitList = append(fruitList[0:3])
	// OUTPUT : [Banana Mango]
	// 3 here is excluding
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 10
	highScores[1] = 0
	highScores[2] = 5
	highScores[3] = 10
	// highScores[4] = 8 This will give us an error

	highScores = append(highScores, 5, 7, 8)
	fmt.Println(highScores)

	sort.Ints(highScores)
	// sort.Strings(fruitList)
	// fmt.Println(fruitList)

	fmt.Println("High Scores after sorting", highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove a value from a slice based on index
	var courses = []string{"javascript", "reactjs", "java", "typescript", "python"}
	fmt.Println(courses)
	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("After removing the course at index 2:", courses)

	// var ans = []string{"Prathamesh"}
	// ans = append(ans, fruitList...)
	// fmt.Println(ans)

	// adding a new value to slice at a particular index
	// var courses1 = []string{}
	// var courses2 = []string{}
	// courses1 = append(courses[:index])
	// courses2 = append(courses[index:])

	// var ans = []string{}
	// ans = append(ans, courses1...)
	// ans = append(ans, "C++")
	// ans = append(ans, courses2...)
	// fmt.Println(ans)

	// var arr = []int{1, 2, 3, 4, 5, 6}
	// fmt.Println(arr)

	// var arr1 = []int{}
	// var arr2 = []int{}

	// arr1 = append(arr1, arr[:2]...)
	// arr2 = append(arr2, arr[2:]...)
	// var newArr = []int{}
	// newArr = append(newArr, arr1...)
	// newArr = append(newArr, 7)
	// newArr = append(newArr, arr2...)
	// fmt.Println(newArr)
}
