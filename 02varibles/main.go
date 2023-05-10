package main

import "fmt"

//  this is public varibale which start with capital letter
const LoginToken string = "badl"

func main() {
	username := "Divyansh Sharma"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	// default value
	var value int
	fmt.Println(value)
	fmt.Printf("Variable is of type: %T \n", value)

	// implicit type
	var website = "divuansh"
	fmt.Println(website)

	fmt.Println(LoginToken)
	fmt.Printf("Varible is of type : %T\n", LoginToken)
}
