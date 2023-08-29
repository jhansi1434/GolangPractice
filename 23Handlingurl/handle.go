package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "heeps://loc.dev:3000/learn?coursename=reactjs&paymentid=ghni456"

func main() {

	fmt.Println(myUrl)

	result,_:=url.Parse(myUrl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	
}
