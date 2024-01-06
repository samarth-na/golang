package resource

import "math/rand"

func GenRandInt(limit int) int {
	randInt := rand.Intn(limit)
	if randInt == 0 {
		randInt = GenRandInt(limit)
	}
	return randInt
}

func GenRandArray(limit int, size int) []int {
	randArray := make([]int, size)
	for i := 0; i < size; i++ {
		randArray[i] = GenRandInt(limit)
	}
	return randArray
}
func GenRand3X3Matrix(integer int) [3][3]int {
	var mat [3][3]int
	for index, value := range mat {
		for index2 := range value {
			mat[index][index2] = GenRandInt(integer)
		}

	}
	return mat
}

func GenRand2X2Matrix(integer int) [2][2]int {
	var mat [2][2]int
	for index, value := range mat {
		for index2 := range value {
			mat[index][index2] = GenRandInt(integer)
		}
	}
	return mat
}
func GenRandSlice(limit int, size int) []int {
	randSlice := make([]int, size)
	for i := 0; i < size; i++ {
		randSlice[i] = GenRandInt(limit)
	}
	return randSlice
}
