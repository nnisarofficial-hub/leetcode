package main

func TwoSum(nums []int, target int) []int {
	size := len(nums)
	for index := range nums {
		for i := index + 1; i < size; i++ {
			if nums[index]+nums[i] == target {
				return []int{index, i}
			}
		}
	}
	return nil
}
