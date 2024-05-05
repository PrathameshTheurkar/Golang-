package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// PerformGetRequest()
	// PerformPostRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)

	checkNilError(err)

	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content Length:", response.ContentLength)

	// Create a Builder
	var responseString strings.Builder

	content, err := io.ReadAll(response.Body)
	checkNilError(err)

	byteCount, err := responseString.Write(content)
	checkNilError(err)

	fmt.Println("Byte Count is:", byteCount)

	fmt.Println(responseString.String())

	// fmt.Println(string(content))

}

func PerformPostRequest() {
	const myurl = "http://localhost:8000/post"

	// fake json payload
	requestBody := strings.NewReader(`
	{
		"coursename" : "Golang",
		"price" : 0
	}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	checkNilError(err)
	defer response.Body.Close()
	// var responseString strings.Builder
	content, err := ioutil.ReadAll(response.Body)
	checkNilError(err)
	// byteCount, err := responseString.Write(content)
	// fmt.Println(byteCount)
	// fmt.Println(responseString.String())

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//form data
	data := url.Values{}
	data.Add("firstName", "Prathamesh")
	data.Add("lastName", "Theurkar")
	data.Add("Email", "prathameshtheurkar037@gmail.com")

	response, err := http.PostForm(myurl, data)
	checkNilError(err)
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	checkNilError(err)

	fmt.Println(string(content))

}

func doPostReq() {
	reqBody := strings.NewReader(`
	{
		"courseName" : "golang",
		"price" : 0
	}
	`)

	res, err := http.Post("http://localhost:8000/post", "application/json", reqBody)
	checkNilError(err)
	defer res.Body.Close()

	databytes, err := io.ReadAll(res.Body)
	checkNilError(err)

	var stringResponse strings.Builder

	byteCount, err := stringResponse.Write(databytes)
	checkNilError(err)

	fmt.Println("Count : ", byteCount)
	fmt.Println("Response : ", stringResponse.String())
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
