package medium

import (
	"fmt"
	"sort"
)

// Link problem:
// https://leetcode.com/problems/find-all-duplicates-in-an-array/

func FindDuplicates(nums []int) []int {
	var output []int
	for i := 0; i < len(nums); i++ {
	innerloop:
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == nums[i] {
				output = append(output, nums[i])
				break innerloop
			}
		}
	}
	sort.Ints(output)
	return output
}

func Main_FindDuplicates() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println("==========FindDuplicates=========")
	fmt.Println("Input\t: ", nums)
	fmt.Println("Output\t: ", FindDuplicates(nums))
	fmt.Println("")
}
