package basics

import (
	"fmt"
)

func recc(index int, limit int, step int) {
	fmt.Println(index)

	newindex := index + step

	if newindex <= limit {
		recc(newindex, limit, step)
	}
}
