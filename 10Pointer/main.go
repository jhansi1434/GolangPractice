package main

import "fmt"

func main() {
	

	var ptr int=67
	var num *int
	fmt.Println("Variable Value:",num)
	fmt.Println(ptr)

	//pointers
	myNumber:=25
	var ptr1 = &myNumber
	fmt.Println("Variable Value:",ptr1)  //it gives address of the variable
	fmt.Println("Variable Value:",*ptr1) //it gives value of the variable
}
