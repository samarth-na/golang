package resourse

import "fmt"

func Printx(mainArr [5][5]int) {
	for _, subArr := range mainArr {
		for _, value := range subArr {
			if value != 0 {
				fmt.Print("x")
			}
			fmt.Print(" ")
		}
		fmt.Printf("\n")
	}
}

func Sortarray() {
	mainArr := [5][5]int{
		{1, 0, 0, 0, 1},
		{0, 1, 0, 1, 0},
		{0, 0, 1, 0, 0},
		{0, 1, 0, 1, 0},
		{1, 0, 0, 0, 1},
	}
	Printx(mainArr)

	// fmt.Println(mainArr[1][2])
}
