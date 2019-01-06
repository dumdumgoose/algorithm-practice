package easy

import "sort"

// this is the solution I came up, it's not very fast because the sorting takes some time
func majorityElement(nums []int) int {
	sort.Ints(nums)
	counter := 0
	majority := len(nums) / 2

	for i, n := range nums {
		if i > 0 {
			if nums[i] == nums[i-1] {
				counter++
			} else {
				counter = 1
			}
		} else {
			counter = 1
		}

		if counter > majority {
			return n
		}
	}

	return 0
}

// this is the one in the solution and shorter and simpler, but the speed is about the same, because sorting takes time
func majorityElementSorting(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// an O(n) time and O(1) space solution found in the discussion, Boyer-Moore Voting Algorithm
func majorityElementFast(nums []int) int {
	majority := nums[0]
	counter := 1
	for i := 1; i < len(nums); i++ {
		if counter == 0 {
			counter++
			majority = nums[i]
		} else if majority == nums[i] {
			counter++
		} else {
			counter--
		}
	}
	return majority
}
