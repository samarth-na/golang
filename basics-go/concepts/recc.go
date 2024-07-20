package basics

import "fmt"

var x = 100

func Recc() {
	if x <= 0 {
		fmt.Print("end")
		return
	}
	x--
	fmt.Println(x)
}
