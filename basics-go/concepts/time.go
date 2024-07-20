package basics

import (
	"fmt"
	"time"
)

func Time() {
	o := 1
	i, j := 1, 2
	timeMark := time.Now()
	fmt.Println(timeMark)
	// fmt.Scan(&i)
	i = 23
	for o <= i {
		j += o
		fmt.Println(j)
		o++
	}
	timeMark2 := time.Now()
	fmt.Println(timeMark)
	fmt.Println(timeMark2)
}

// type i int
// func (int i) Time() {
// 	fmt.Print("i is", int)
// 	fmt.Print(time.Now())
// }
