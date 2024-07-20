package main

import (
	"alderaan/resource"
	"fmt"
)

func takeInp() {
	var numOfArray, rangeOfnum int

	fmt.Println("enter the number of arrays")
	fmt.Scan(&numOfArray)
	fmt.Println("number of number in array")
	fmt.Scan(&rangeOfnum)
}

func main() {
	resource.GenRandInt(100)
	takeInp()
}
