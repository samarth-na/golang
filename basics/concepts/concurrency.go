package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func concur(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func Concurrency() {
	for n := 0; n < 10; n++ {
		go concur(n)
	}
	var input string
	fmt.Scanln(&input)
}
