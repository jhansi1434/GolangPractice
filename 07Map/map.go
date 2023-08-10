package main

import "fmt"

func main() {
	courses := make(map[string]string)

	courses["TS"] = "TypeScript"
	courses["JS"] = "JavaScript"
	courses["GO"] = "go"
	courses["PN"] = "python"

	fmt.Println(courses)
	fmt.Println("js shorts for:",courses["JS"])
     
	delete(courses,"PN")
	fmt.Println("deleted courses:",courses)
}
