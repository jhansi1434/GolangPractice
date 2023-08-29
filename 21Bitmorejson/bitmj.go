package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string  `json:"coursename"`
	Price    int
	Platform string  `json:"website"`
	Password string   `json:"-"`
	Tags     []string  `json:"tags,omitempty"`
}

func main() {
// EncodeJson()
consumeData()
}

// the main purpose of this code is to get actual consumable JSON data..by using type and structs.
//In this using alias names also


func EncodeJson() {
	 
	lcoCourses := []course{
		{"ReactJs Bootcamp",296,"learnonline.go","jhansi143",[]string{"web-dev", "js"}},

		{
			"ReactJs Bootcamp",
			296,
			"learnonline.go",
			"jhansi143",
			[]string{"web-dev", "js"}},

		{
			"ReactJs Bootcamp",
			296,
			"learnonline.go",
			"jhansi143",
			nil},
	}

	finaljson, err := json.MarshalIndent(lcoCourses, "","\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finaljson)
}


func consumeData(){

  jsonFormat :=[]byte(`
  {
	"coursename": "ReactJs Bootcamp",
	"Price": 296,
	"website": "learnonline.go",
	"tags": ["web-dev","js"]
}
  `)

  var lcoCourse course

  checkValid := json.Valid(jsonFormat)
if checkValid{
	fmt.Println("Json is valid")
	json.Unmarshal(jsonFormat, &lcoCourse)
	fmt.Printf("%#v\n",lcoCourse)
}else{
	fmt.Println("Json was Not valid")
}	 
  

//some case where you want to just add data tokey value

var mydata map[string]interface{}
json.Unmarshal(jsonFormat,&mydata)
fmt.Printf("%#v\n",mydata)

for k,v:=range mydata{
	fmt.Printf("key is %v and value is %v and Type is %T\n",k,v,v)
}

}
