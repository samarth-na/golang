package basics

import "fmt"

func add() {
	var a, b int
	fmt.Print("no1:")
	fmt.Scan(&a)

	fmt.Print("no2:")
	fmt.Scan(&b)

	fmt.Println(a + b)
}
