package main

import "fmt"

func IntToChr(integer int) {
	for integer >= 1 {
		singleint := integer % 10
		integer /= 10
		fmt.Println(singleint)

	}
}
