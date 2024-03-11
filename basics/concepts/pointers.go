package basics

import "fmt"

func Pointer(a int) int {
	b := &a
	fmt.Println(*b)
	return *b
}
