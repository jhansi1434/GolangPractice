package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"net/url"
	
)

func main() {
	// Data()
	// postData()
	formData()
}

func Data() {
	const url = "http://localhost:8000/get"
	res, err := http.Get(url)

	errorMes(err)

	defer res.Body.Close()
	fmt.Println("status code", res.Status)
	fmt.Println("content length is", res.ContentLength)

	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
}

func postData() {
	const url = "http://localhost:8000/post"

	jsonData := strings.NewReader(`
		{
			"name":"jhansi",
			"price":0,
			"platform":"learncode.in"
		}
	`)

	res, err := http.Post(url, "application/json", jsonData)
	errorMes(err)

	defer res.Body.Close()
	fmt.Println("status code", res.Status)
	fmt.Println("content length is", res.ContentLength)

	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))

}

func formData() {
	const myUrl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstname", "jhansi")
	data.Add("lastname", "pasupuleti")
	data.Add("email", "jhansi@email.com")

	res, err := http.PostForm(myUrl, data)
	errorMes(err)
	defer res.Body.Close()
	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
}

func errorMes(err error) {
	if err != nil {
		panic(err)
	}
}
