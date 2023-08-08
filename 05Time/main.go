package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")
	presentTime:=time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 monday"))
	//this is another method to get exact format for present time,date,year(01-02-2006 15:04:05 monday)

	createDate:=time.Date(2020,time.February,10,23,0,0,0,time.UTC)
	fmt.Println(createDate.Format("01-02-2006 15:04:05 monday"))
}
