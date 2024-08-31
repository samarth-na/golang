package exercise

import (
	"fmt"
)

func IntToChr(integer int) int {
	if integer < 10 {
		return integer
	}
	fmt.Println(integer % 10)
	return IntToChr(integer / 10)
}
