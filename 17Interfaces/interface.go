// package main

// import "fmt"

// type alldetails interface {
// 	Details()
// 	Details1()
// }
// type Studetails struct {
// 	name string
// 	age  int
// 	year int
// }

// type Empldetails struct {
// 	email  string
// 	mobile int
// 	Salary int
// }

// func (s Studetails) Details() {
// 	fmt.Println("student details", s.name, s.age)
// }

// func (s Studetails) Details1() {
// 	fmt.Println(s.year)
// }

// func (e Empldetails) Details() {
// 	fmt.Println("employee details", e.email, e.mobile)
// }

// func (e Empldetails) Details1() {
// 	fmt.Println(e.Salary)
// }

// func de(d alldetails) {
// 	fmt.Println(d)
// 	d.Details()
// 	d.Details1()
// }

// func main() {

// 	n := Studetails{
// 		name: "jhansi",
// 		age:  23,
// 		year: 2000,
// 	}
// 	e := Empldetails{
// 		email:  "jhansi@gmail.com",
// 		mobile: 989887655,
// 		Salary: 20000,
// 	}
// 	de(n)
// 	de(e)
// }

// package main

// import (
//     "fmt"
//     "math"
// )

// type geometry interface {
//     area() float64
//     perim() float64
// }

// type rect struct {
//     width, height float64
// }
// type circle struct {
//     radius float64
// }

// func (r rect) area() float64 {
//     return r.width * r.height
// }
// func (r rect) perim() float64 {
//     return 2*r.width + 2*r.height
// }

// func (c circle) area() float64 {
//     return math.Pi * c.radius * c.radius
// }
// func (c circle) perim() float64 {
//     return 2 * math.Pi * c.radius
// }

// func measure(g geometry) {
//     fmt.Println(g)
//     fmt.Println(g.area())
//     fmt.Println(g.perim())
// }

// func main() {
//     r := rect{width: 3, height: 4}
//     c := circle{radius: 5}

//     measure(r)
//     measure(c)
// }

package main

import "fmt"

type SalaryCaluculater interface {
	CaluculateSalary() int
}
type Permanent struct {
	empId    int
	basicPay int
	pf       int
}

type Contract struct {
	empId    int
	basicPay int
}

func (p Permanent) CaluculateSalary() int {
	return p.basicPay + p.pf
}

func (c Contract) CaluculateSalary() int {
	return c.basicPay
}

func total(s []SalaryCaluculater) {
	expense := 0
	for _, v := range s {
		expense = expense + v. CaluculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

func main() {
	pemp1 := Permanent{
        empId:    1,
        basicPay: 5000,
        pf:       20,
    }
    pemp2 := Permanent{
        empId:    2,
        basicPay: 6000,
        pf:       30,
    }
    cemp1 := Contract{
        empId:    3,
        basicPay: 3000,
    }
    employees := []SalaryCaluculater{pemp1, pemp2, cemp1}
    total(employees)
}
