package easy

import "fmt"

// link:
// https://leetcode.com/problems/two-sum/

func TwoSum(nums []int, target int) []int {
	var a []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				a = append(a, i)
				a = append(a, j)
			}
		}
	}
	return a
}

func Main_TwoSum() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println("==============TwoSum=============")
	fmt.Println("Input\t: nums = ", nums, "& target = ", target)
	fmt.Println("Output\t: ", TwoSum(nums, target))
	fmt.Println("")
}
