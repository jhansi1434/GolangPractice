package main

import "fmt"

//buffered channel

// func main() {
// 	ch := make(chan int, 1)
// 	ch <- 89
// 	val := <-ch
// 	fmt.Println(val)
// 	unbuffered()
// 	rw()
// }

//unbuffered channel
//in this type we get "all goroutines are asleep - deadlock this error"

// func unbuffered() {
// 	ch := make(chan int)
// 	ch <- 89
// 	val := <-ch
// 	fmt.Println(val)
// }

//how to read and write to a go channel

func main() {
	ch := make(chan int)

	//write to the channel

	go func() {
		ch <- 42
	}()

	//read to the channel
	
	value := <-ch
	fmt.Println(value)

	//write to the channel again

	go func() {
		ch <- 45
	}()

	//read to the channel again

	value = <-ch
	fmt.Println(value)
}
