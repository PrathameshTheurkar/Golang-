package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string
	Price    int
	Platform string
	Password string
	Tags     []string
}

func main() {
	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"Reactjs", 1000, "lco", "12345", []string{"web-dev", "reactjs", "js"}},
		{"Mern Bootcamp", 1000, "lco", "1245", []string{"web-dev", "fullstack", "js"}},
		{"Golang", 1000, "lco", "123s", nil},
	}

	// package this data as JSON data

	// finalJson, err := json.MarshalIndent(lcoCourses, "lco", "\t")
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	checkNilError(err)
	fmt.Printf("%s\n", finalJson)

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
