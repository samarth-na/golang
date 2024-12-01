package resourse

import "fmt"

func By5by2into5(input int) int {
	input = (input / 5) / 2
	fmt.Println(input * 5)
	return 0
}
