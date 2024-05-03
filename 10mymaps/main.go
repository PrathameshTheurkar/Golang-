package main

import "fmt"

func main() {

	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"

	fmt.Println(languages)
	fmt.Println(languages["js"])

	delete(languages, "rb")
	fmt.Println(languages)

	// loops in map
	// for _, value := range languages {
	// 	fmt.Printf("For key v the value is %v\n", value)
	// 	// fmt.Println(key, value)
	// }
	for key, value := range languages {
		fmt.Printf("For key %v the value is %v\n", key, value)
		// fmt.Println(key, value)
	}

}
