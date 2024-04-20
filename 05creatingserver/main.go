package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video - My APIs")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:1111/get"
	resp, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Status code: ", resp.StatusCode)
	fmt.Println("Content length", resp.ContentLength)
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(content))
	byteCount, _ := responseString.Write(content)
	fmt.Println("Bytecount is: ", byteCount)
	fmt.Println(responseString.String())
}
