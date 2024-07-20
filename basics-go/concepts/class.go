package basics

import (
	"fmt"
)

// Define the Address struct with explicit field types
type Address struct {
	City    string
	Country string
	Pincode int
}

// Define the Person struct with explicit field types and pointer to Address

type Person struct {
	Name       string
	Age        int
	Address    *Address
	WeightInKg float64
	HeightInM  float64
	Married    bool
}

var list [10]int

// Calculate BMI method using explicit receiver and return type
func (p *Person) Bmi() float32 {
	bmi32 := float32(p.WeightInKg / (p.HeightInM * p.HeightInM))
	// bmi32 := float32(bmi)
	return bmi32
}

func classmain() {
	// Create Address instance

	mattAddress := Address{
		City:    "London",
		Country: "UK",
		Pincode: 12345,
	}

	// Create Person instance with pointer to Address
	matt := Person{
		Name:       "Matt",
		Age:        30,
		Address:    &mattAddress, // Reference the created address
		WeightInKg: 68.5,
		HeightInM:  1.55,
		Married:    true,
	}
	{
		fmt.Printf("%+v\n", matt) // Raw output
	}

	// Formatted output
	fmt.Printf("%s is %d years old\n", matt.Name, matt.Age)
	fmt.Printf("%s is in %s\n", matt.Address.City, matt.Address.Country)
	fmt.Printf("%f kg and %f m, BMI is: %.2f\n", matt.WeightInKg, matt.HeightInM, matt.Bmi())

	if matt.Married {
		fmt.Printf("%s is married\n", matt.Name)
	} else {
		fmt.Printf("%s is single\n", matt.Name)
	}
}

// end file
