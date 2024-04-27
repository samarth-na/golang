package basics

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}

func twoSum(nums []int, target int) []int {
	in := func(i int, j int) (int, int) {
		if nums[i]+nums[j] == target {
			return i, j
		}
		if j < len(nums) {
			j++
		}
		if i < len(nums) {
			i++
		}
		return i, j
	}
	fmt.Println(in(0, 1))

	i, j := in(0, 1)

	anslist := []int{i, j}
	return anslist
}
