package main

import "fmt"

func main() {
	fmt.Println("welcome to ifelse")
	var result string
	Taste:="excelent"

	if Taste=="super"{
        result="yes it is super"
	}else{
		result="NO it is super "
	}
	fmt.Println(result)

	var age string
	Age:=10
	if Age > 20{
		age="my is not 20"
	} else if Age < 9{
		age="my age is 23"
	} else{
		age="my age is above 20"
	}

	fmt.Println(age)

	if 9%2==0{
		fmt.Println("Number is even")

	}else{
		fmt.Println("Number is odd")
	}


	if Salary:=2000;Salary<3000{
		fmt.Println("my salary is 2000")
	}else {
		fmt.Println("my salary is not 1000 ")
	}
}
