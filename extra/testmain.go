package main

import "fmt"

func ret() (int, int) {
	x := 2
	return x, 0

}

func nomain() {
	x, _ := ret()

	fmt.Println(x)

}
