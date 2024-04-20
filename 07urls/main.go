package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=gfhhj747"

func main() {
	fmt.Println("Welcome to urls in golang")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Println(qparams)
	fmt.Println(qparams["coursename"])

	for key, val := range qparams {
		fmt.Println("Key is: ", key)
		fmt.Println("Value is: ", val)
	}

	// constructing a url

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
