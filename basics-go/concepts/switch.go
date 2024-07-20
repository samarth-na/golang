package basics

import "fmt"

func Switch() {
	value := "B"

	switch value {
	case "A":
		fmt.Println("Value is A")
	case "B":
		fmt.Println("Value is B")
	default:
		fmt.Println("Value is not A or B")
	}
}
