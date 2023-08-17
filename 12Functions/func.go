package main

import "fmt"

func main() {
	fmt.Println("welcome to functions in go lang")
greeting()
  result:=addition(2,3)
  fmt.Println(result)

  pro:=proAdd(2,3,4,5)
  fmt.Println("value is :",pro)

}
 
func greeting(){
	fmt.Println("Welcome to greeting ")
}

//addition of two numbers

func addition(valueOne int,valueTwo int)int{
return valueOne+valueTwo
  
}

//addition in pro 

func proAdd(value ...int)int  {
	total:=0
	for _,val:=range value{
		total+=val
		fmt.Println(val)
	}
	return total

}