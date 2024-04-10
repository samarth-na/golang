package exercise

import "fmt"

var birthdate, birthmonth, birthyear, currentdate, currentmonth, currentyearr int

var totalyear, totalmonth, totaldays int

func dataIn() {
	fmt.Print("enter your birth date- ")
	fmt.Scan(&birthdate)
	for birthdate > 31 {
		fmt.Println("you must have made any mistake you cant have date above 31 ")
		fmt.Print("try again :-")
		fmt.Scan(&birthdate)
	}
	fmt.Print("enter your birth month- ")
	fmt.Scan(&birthmonth)
	for birthmonth > 12 {
		fmt.Println("you must have made any mistake you cant have month above 12")
		fmt.Print("try again :-")
		fmt.Scan(&birthmonth)
	}
	fmt.Print("enter your birth year- ")
	fmt.Scan(&birthyear)
	for birthyear < 100 {
		fmt.Println("enter in full YYYY notion like 1999")
		fmt.Print("try again :-")
		fmt.Scan(&birthyear)
	}
	fmt.Print("enter current date- ")
	fmt.Scan(&currentdate)
	for currentdate > 31 {
		fmt.Println("you must have made any mistake you cant have date above 31 ")
		fmt.Print("try again :-")
		fmt.Scan(&currentdate)
	}
	fmt.Print("enter current month- ")
	fmt.Scan(&currentmonth)
	for currentmonth > 12 {
		fmt.Println("you must have made any mistake you cant have month above 12")
		fmt.Print("try again :-")
		fmt.Scan(&currentmonth)
	}
	fmt.Print("enter current year- ")
	fmt.Scan(&currentyearr)
	for currentyearr < 1000 {
		fmt.Println("enter in full YYYY notion like 1999")
		fmt.Print("try again :-")
		fmt.Scan(&currentyearr)
	}
	for currentyearr < birthyear {

		fmt.Println("today can be before you were born")
		break
	}
}

func findingage() {
	totalyear = currentyearr - birthyear

	if currentmonth == birthmonth && currentdate < birthdate {
		currentdate += 30
		totalmonth = 11
		totalyear--
		totaldays = currentdate - birthdate
	} else if currentmonth < birthmonth {
		currentmonth = currentmonth + 12
		totalyear--
	}
	totalmonth = currentmonth - birthmonth
	if currentdate < birthdate {
		totalmonth--
		currentdate += 30
		totaldays = currentdate - birthdate
	}

	totaldays = currentdate - birthdate
}

func Age() {
	dataIn()
	findingage()
	fmt.Println("you are", totaldays, "days", totalmonth, "months and", totalyear, "years old")
}
