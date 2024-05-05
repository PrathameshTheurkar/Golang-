package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`              // '-' this doesn't shows this field
	Tags     []string `json:"tags,omitempty"` //'omitempty' this just omit the field if it is empty
}

func main() {
	EncodeJson()

	// DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"Reactjs", 1000, "lco", "12345", []string{"web-dev", "reactjs", "js"}},
		{"Mern Bootcamp", 1000, "lco", "1245", []string{"web-dev", "fullstack", "js"}},
		{"Golang", 1000, "lco", "123s", nil},
	}

	// package this data as JSON data

	// finalJson, err := json.Marshal(lcoCourses)
	// finalJson, err := json.MarshalIndent(lcoCourses, "lco", "\t")
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	checkNilError(err)
	fmt.Printf("%s\n", finalJson) //%s format specifier is used to specify the string
	//whereas %v format specifier is used to print the value of variable in default format

	// DECODING THE JSON / CONSUMING THE JSON

	fmt.Printf("Type of finalJson is %T\n", finalJson)

	checkValidJson := json.Valid(finalJson)

	if checkValidJson {
		fmt.Println("Json is valid")
		var myOnlineData []map[string]interface{}
		json.Unmarshal(finalJson, &myOnlineData)
		fmt.Printf("%#v\n", myOnlineData)
	} else {
		fmt.Println("Json is not valid")
	}

}

func DecodeJson() {
	jsonFromWeb := []byte(`
	{
		"coursename": "Reactjs",
		"price": 1000,
		"platform": "lco",
		"tags": ["web-dev","reactjs","js"]
	}
	`)

	var lcocourse course

	checkValidJson := json.Valid(jsonFromWeb)

	if checkValidJson {
		fmt.Println("Json is valid")
		json.Unmarshal(jsonFromWeb, &lcocourse)
		fmt.Printf("%#v\n", lcocourse) //%v format specifier is used to print the value of variable in default format
		fmt.Printf("%v\n", lcocourse.Name)
	} else {
		fmt.Println("Json is not valid")
	}

	// some other cases where you just want to add data to key value

	fmt.Println("some other cases where you just want to add data to key value")
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("Key : %v & value : %v\n", key, value)
	}

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
