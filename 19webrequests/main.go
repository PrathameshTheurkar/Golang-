package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// const url = "https://lco.dev"
const url = "http://localhost:8000/get"

func main() {
	response, err := http.Get(url)
	checkNilError(err) //handle the error
	fmt.Printf("Response is of type %T\n", response)
	defer response.Body.Close() //caller's responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)
	checkNilError(err)
	content := string(databytes)
	fmt.Println(content)

}

func checkNilError(err error) {
	if err != nil {
		fmt.Println("This is error")
		panic(err)
	}
}
