package main

import "fmt"

func main() {
	//it is a last in first out //
	defer fmt.Println("world")
	defer fmt.Println("hello")
	fmt.Println("welcome to defer")
	deferfun()
}
func deferfun(){
	for d:=0;d<5;d++{
		defer fmt.Println(d)
	}
}
