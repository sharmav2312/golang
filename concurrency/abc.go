package main

import (
	"fmt"
)

func Abc() {
	fmt.Println("abc")

	i := []int{1, 5, 6}

	for a := range i {
		fmt.Println(i[a])
	}
}
