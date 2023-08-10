package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitsList1 = []string{"Apple", "mango", "peach"}
	fruitsList1 = append(fruitsList1, "jackfruits")
	fmt.Println(fruitsList1)

	fruitsList1 = append(fruitsList1[1:3])
	//{:1}it gives only first value
	//{1:3}it gives 1st and 2nd value
	fmt.Println(fruitsList1)

	names := make([]string, 3)
	names[0] = "jhansi"
	names[1] = "jhanu"
	names[2] = "chithi"
	names = append(names, "harish", "rani")
	fmt.Println(names)

	//some more these are using in slices only
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.StringsAreSorted(names))

	//how to remove a value from slices based on index

   var Age=[]int{23,24,75,26,43,56}
   index:=2
   fmt.Println(Age)
   Age=append(Age[:index],Age[index+1:]...)
   fmt.Println(Age)

}
