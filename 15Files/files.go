package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	content := "this file is in go lang content klklkl"
	file, err := os.Create("./myLocgofile.txt")
	checkNilerr(err)

	length, err := io.WriteString(file, content)
	checkNilerr(err)

	fmt.Println("length is:", length)
	file.Close()
	readFile("./mylocgofile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilerr(err)
	fmt.Println("Text data inside the file is \n", databyte)
}

func checkNilerr(err error) {
	if err != nil {
		panic(err)
	}
}
