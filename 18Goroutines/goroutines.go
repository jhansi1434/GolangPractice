package main

import (
	"fmt"
	"net/http"
	// "time"
)

func main() {
//go rotuines are used just adding "go" keyword

//    go greeter("Hello")
//    greeter("world")
 
webSiteList:=[]string{
"https://lco.dev",
"https://go.dev",
"https://google.com",
"https://fb.com",
"https://github.com",
}

for _,web:=range webSiteList{
	getStatusCode(web)
}



}
// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(1*time.Second)
// 		fmt.Println(s)
// 	}	
// }


func getStatusCode(endpoint string){
res,err:=http.Get(endpoint)
if err!=nil{
	fmt.Println("OOPS in endpoint ")
}else{
	fmt.Printf("%d status code for %s\n",res.StatusCode,endpoint)
}

}
