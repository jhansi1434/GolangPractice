package main

import "fmt"

func main() {
	fmt.Println("welcome to loop in go lang")

	//months:=[]string{"jan","feb","march","april","may","june...etc"}

	
	// for d:=0; d<len(months);d++{
	// 	fmt.Println(months[d])
	// }

	//another methods

	// for i:=range(months){
	// 	fmt.Println(months[i])
	// }

	// for index,months:=range months{
	// 	fmt.Println("index is", index ,"and its value is",months)
	// }

	value:=1

	for value<10{
      
if value==2{
	goto lco
	break
}

       if value==5{
		value++
		break
		
	   }

		fmt.Println("value is:",value)
		value++

		
	}
	lco:
		fmt.Println("jumping into learncode.online")
	
}
