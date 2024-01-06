package basics

import (
	"fmt"
	"time"
)

var t int

func Time() {
	o := 1
	var i, j int = 1, 2
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
