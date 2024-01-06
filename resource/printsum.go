package resource

import "fmt"

func printsum() {
	sum := func(a int, b int) int {
		return a + b
	}
	var val1, val2 int
	fmt.Scanf("%v", &val1)
	fmt.Scanf("%v", &val2)
	total := sum(val1, val2)
	fmt.Println(total) // Output: 30
}
