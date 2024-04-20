package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("new web request")

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("response is of type: %T\n", resp)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	content := string(body)

	fmt.Println(content)
}
