package exercise

import "fmt"

func IntToChr(integer int) []int {
	newList := []int{}
	for integer >= 1 {
		singleDigit := integer % 10
		integer /= 10
		fmt.Println(singleDigit)
		newList = append(newList, singleDigit)
	}
	return newList
}
