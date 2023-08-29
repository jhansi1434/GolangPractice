package main

import (
	"fmt"
	"io"
	"net/http"
)

const url string = "https://lco.dev"

func main() {
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is type%T\n", res)

	defer res.Body.Close()

	databytes,err:=io.ReadAll(res.Body)

	if err!=nil{
		panic(err)
	}
	
    content:=string(databytes)

	fmt.Println(string(content))
}
