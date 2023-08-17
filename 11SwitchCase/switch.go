package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch and case")

	rand.Seed(time.Now().UnixNano())
	diceNumber:=rand.Intn(6)+1
     
	fmt.Println("value of dice number:",diceNumber)

	switch diceNumber{
	case 1:
		fmt.Println("dice value is 1 you can open")
	case 2:
		fmt.Println("dice value 2")
		fallthrough
	case 3:
		fmt.Println("dice value 3")
		
	case 4:
		fmt.Println("dice value 4")
	case 5:
		fmt.Println("dice value 5")
	case 6:
		fmt.Println("dice value 6 and roll once agin!")
	default:
		fmt.Println("what was that!")

	}
}
