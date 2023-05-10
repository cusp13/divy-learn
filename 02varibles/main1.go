package main

import "fmt"

//  this is public varibale which start with capital letter
const LoginToken string = "badl"

func main() {
	username := "Divyansh Sharma"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var smallVal float32 = 25.7656352
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	fmt.Println(LoginToken)
	fmt.Printf("Varible is of type : %T\n", LoginToken)
}
