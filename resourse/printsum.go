package resourse

import "fmt"

func Printsum() {
	sum := func(a int, b int) int {
		return a + b
	}
	var val1, val2 int
	fmt.Print("enter first number :")
	fmt.Scanf("%v", &val1)
	fmt.Print("\nenter second number :")
	fmt.Scanf("%v", &val2)
	total := sum(val1, val2)
	fmt.Println(total) // Output: 30

	Pattern()
}
