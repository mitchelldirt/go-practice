package main

// func twoSum(nums []int, target int) []int {
// 	// if there are only two numbers then return first two indexes
// 	if (len(nums) == 2) {
// 		return []int{0,1}
// 	}

// 	// Lo op over the remaining numbers to find the pair that equals the target
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if (nums[i] + nums[j] == target) {
// 				return []int{i, j}
// 			}
// 		}
// 	}

// 	return nil
// }