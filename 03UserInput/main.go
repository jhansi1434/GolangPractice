package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "to the go lang"
	fmt.Println(welcome)

	reader :=bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating of my app")

	//comma ok||err err

	input, _:=reader.ReadString('\n')
	fmt.Println("thanks for rating",input)
	fmt.Printf("type of rating %T",input)
}
