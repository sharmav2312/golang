package main

import "fmt"

var jwtToken = 300000  // allowed
const LoginToken = 300 // constant, first letter capital indicates that it is public and it is accessible to any function within this program

// var jwtTok := 3470 ----- not allowed outside methods

func main() {
	var username string = "Vishu"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallValue int = 10000
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T \n", smallValue)

	// float64 has more precision than float32
	var smallFloat float64 = 255.454564646
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and aliases
	var anotherVar int16
	fmt.Println(anotherVar)
	fmt.Printf("another variable is : %T \n", anotherVar)
	// implicit type
	var x = "Vishu"
	fmt.Println(x)

	// no var style, using wallrus operator
	// wallrus operator is allowed only inside the methods

	numberOfUsers := "Pata nahi"
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Login Token type  : %T \n", LoginToken)
}
