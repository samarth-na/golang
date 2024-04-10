package datastructures

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Declare a method for the Person struct
func (p Person) Greet() string {
	return "Hello, my name is " + p.Name
}

func main() {
	// Create an instance of the Person struct
	person := Person{Name: "Alice", Age: 30}

	// Call the Greet method
	greeting := person.Greet()
	fmt.Println(greeting) // Output: Hello, my name is Alice
}
