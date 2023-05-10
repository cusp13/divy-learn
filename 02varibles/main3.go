package main

import "fmt"

//  this is public varibale which start with capital letter
const LoginToken string = "badl"

func main() {
	username := "Divyansh Sharma"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)
	fmt.Println(LoginToken)
	fmt.Printf("Varible is of type : %T\n", LoginToken)
}
