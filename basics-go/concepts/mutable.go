package basics

import "fmt"

func mmain() {
	// Mutable data types
	var arr [5]int
	var slice []int
	var map1 map[string]int

	// Initialize mutable data types
	for i := 0; i < 5; i++ {
		arr[i] = i
	}
	slice = append(slice, 1, 2, 3, 4, 5)
	map1 = make(map[string]int)
	map1["one"] = 1
	map1["two"] = 2
	map1["three"] = 3

	// Immutable data types
	var str string = "Hello, World!"
	var ptr *int = new(int)
	var bool1 bool = true
	var int1 int = 10
	var float1 float64 = 3.14

	// Print mutable and immutable data types
	fmt.Println("Mutable data types:")
	fmt.Println("Array:", arr)
	fmt.Println("Slice:", slice)
	fmt.Println("Map:", map1)
	fmt.Println("Immutable data types:")
	fmt.Println("String:", str)
	fmt.Println("Pointer:", ptr)
	fmt.Println("Boolean:", bool1)
	fmt.Println("Integer:", int1)
	fmt.Println("Float:", float1)
}
