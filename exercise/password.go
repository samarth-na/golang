package exercise

import (
	"fmt"
	"math/rand"
)

var (
	length   int
	letters  = "abcdefghijklmnopqrstuvwxyz"
	capitals = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers  = "0123456789"
	symbols  = "!@#$%^&*()"
)

func chooseChars() string {
	var fornum, forcap, forsym string
	var chars string

	fmt.Println("Do you want capital letters? Type 'y' if yes")
	fmt.Scanln(&forcap)
	fmt.Println("Do you want numbers? Type 'y' if yes")
	fmt.Scanln(&fornum)
	fmt.Println("Do you want symbols? Type 'y' if yes")
	fmt.Scanln(&forsym)

	chars += letters

	if forcap == "y" {
		chars += capitals
	}
	if fornum == "y" {
		chars += numbers
	}
	if forsym == "y" {
		chars += symbols
	}
	return chars
}

func GenPassword() string {
	chars := chooseChars()
	var password string

	fmt.Println("enter the length")
	fmt.Scan(&length)
	fmt.Println(length)

	for len(password) <= length {
		index := len(chars)
		index = rand.Intn(index)
		randomchr := chars[index]
		password += string(randomchr)

	}
	return password
}

func main() {
	fmt.Print(GenPassword())

}
