package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=gjfh4545jnk"

func main() {
	fmt.Println("welcome to handling url in golang")
	fmt.Println(myurl)

	//parsing

	result, _ := url.Parse(myurl)
	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Path:", result.Path)
	fmt.Println("Port:", result.Port())
	fmt.Println("RawQuery:", result.RawQuery)

	qparam := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparam)
	fmt.Println(qparam["coursename"])

	for _, val := range qparam {
		fmt.Println("Param is:", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutes",
		RawPath: "user=hitesh",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
