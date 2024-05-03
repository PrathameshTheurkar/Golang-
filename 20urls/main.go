package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=asd23sdfg"

func main() {
	fmt.Println(myurl)

	// parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	// OUTPUT : https
	fmt.Println(result.Host)
	// OUTPUT: lco.dev:3000
	fmt.Println(result.Path)
	// OUTPUT : /learn
	fmt.Println(result.Port())
	// OUTPUT : 3000
	fmt.Println(result.RawQuery)
	// OUTPUT : coursename=reactjs&paymentid=asd23sdfg

	qparams := result.Query()
	fmt.Printf("The type of query params is %T\n", qparams)

	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])

	for key, val := range qparams {
		fmt.Printf("Key is %v and value is %v\n", key, val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "localhost:3000",
		Path:     "learn",
		RawQuery: "username=prathamesh",
	}

	anotherUrl := partsOfUrl.String() //another way of converting it to string

	fmt.Println(anotherUrl)

	anoURL, _ := url.Parse(anotherUrl)
	fmt.Println("Hostname : ", anoURL.Hostname(), "\nPort : ", anoURL.Port())

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
