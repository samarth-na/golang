package main

import (
	"fmt"
	"strings"
)

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func main() {
	test("Lane,", " happy birthday!")
	test("Elon,", " hope that Tesla thing works out")
	test("Go", " is fantastic")
	fmt.Print(strings.Contains("quhwequhfouq", "ouq"))
	getMonthlyPrice()
}

func test(s1, s2 string) {
	fmt.Println(concat(s1, s2))
}

func getMonthlyPrice() int {
	return 0
}
