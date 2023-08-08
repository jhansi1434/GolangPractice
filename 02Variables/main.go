package main

import "fmt"

func main() {
	var Username string="jhansi"
	fmt.Println(Username )
	fmt.Printf("the usename is type:%T \n" ,Username)

	var isLoggedIn bool=true
	fmt.Println(isLoggedIn)
	fmt.Printf("the usename is type:%T \n" ,isLoggedIn)

	
	var smallVal uint16=250
	fmt.Println(smallVal)
	fmt.Printf("the usename is type:%T \n" ,smallVal)

	var smallFloat float32=2567.898989
	fmt.Println(smallFloat)
	fmt.Printf("the usename is type:%T \n" ,smallFloat)

	//deafault are some aliases
	var anothervariable int32
	fmt.Println( anothervariable)
	fmt.Printf("the usename is type:%T \n" , anothervariable)

	//implicit type
	var website="jhansi"
	fmt.Println(website)

	//no var style
	numberOfUser := 300000
	fmt.Println(numberOfUser)
	
}
