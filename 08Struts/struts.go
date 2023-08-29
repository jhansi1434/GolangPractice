package main

import "fmt"

func main() {
	// fmt.Println("Welcome to struts")
    // details1:= Details{"jhansi",23,20000,true,"jhansi@gmail.com"}

	// fmt.Println(details1)

     type Details struct{
		Name string
		Age int
		Salary int
		status bool
		email string
	 }

	 fmt.Println("Welcome to struts")
    details1:= Details{"jhansi",23,20000,true,"jhansi@gmail.com"}

	fmt.Println(details1)
    fmt.Printf("jhansi details: %+v\n",details1)
    fmt.Printf("name is %v email is %v",details1.Name,details1.email)

}
