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
	fmt.Println("Welcome to JSON creation")
	encodeJson()
}

//encoding of JSON
func encodeJson() {
	lcocourse := []course{
		{"ReactJS Bootcamp", 100, "lco.dev", "abc123", []string{"web-dev", "js"}},
		{"Mern Bootcamp", 199, "lco.dev", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "lco.dev", "hit123", nil},
	}

	//package this data as json data

	finalJson, err := json.MarshalIndent(lcocourse, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}
