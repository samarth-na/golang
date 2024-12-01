package exercise

import (
	"alderaan/resourse"
	"fmt"
)

var (
	secretNum int
	guess     int
)

func Check() {
	var guess int
	fmt.Printf("guess an number between 1 and 100 :-")
	ans := resourse.GenRandInt(100)
	for guess != ans {
		fmt.Scanf("%d", &guess)
		if guess > ans {
			fmt.Println("youre wrong the number is smaller")
		}
		if guess < ans {
			fmt.Println("youre wrong the number is bigger")
		}
		if guess == ans {
			fmt.Println("yea youre right the number was", ans)
		}

	}
}
