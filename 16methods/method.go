// package main

// import "fmt"

// type Details struct{
// 	Name string
// 	Age int
// 	Salary int
// 	status bool
// 	email string
//  }

// func main() {
// 	// fmt.Println("Welcome to struts")
//     // details1:= Details{"jhansi",23,20000,true,"jhansi@gmail.com"}

// 	// fmt.Println(details1)

// 	 fmt.Println("Welcome to struts")
//     details1:= Details{"jhansi",23,20000,true,"jhansi@gmail.com"}

// 	fmt.Println(details1)
//     fmt.Printf("jhansi details: %+v\n",details1)
//     fmt.Printf("name is %v email is %v\n",details1.Name,details1.email)
//      details1.getStatus()
//      details1.newEmail()
// }

// func (u Details)getStatus(){
// 	fmt.Println("status is:",u.status)
// }

// func (u Details)newEmail(){
// 	u.email="jhansinew@email.com"
// 	fmt.Println("my new email is:",u.email)
// }

package main

import "fmt"

type Studetails struct {
	Name string
	Age  int
}

type Empldetails struct {
	email  string
	mobile int
}

func (s Studetails) Details() {
	fmt.Println("student details are :", s.Name, s.Age)
}

func (e Empldetails) Details() {
	fmt.Println("employe detials are:", e.email, e.mobile)
}

func main() {
	d := Studetails{
		Name: "jhansi",
		Age:  23,
	}
	d.Details()
	// fmt.Println("student details ",d.Details())
	E := Empldetails{
		email:  "jhansi@email.com",
		mobile: 9490790529,
	}
	E.Details()
	// fmt.Println("employe details", E.Details())
}
