package main

import "fmt"

func main() {
	input := []int{1, 1, 0, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(input))
}

// https://leetcode.com/explore/learn/card/fun-with-arrays/521/introduction/3238/

func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	for i := 0; i < len(nums); i++ {
		count := 0
		if nums[i] == 1 {
			count++
			for j := i + 1; j < len(nums); j++ {
				if nums[j] == 1 {
					count++
				} else {
					break
				}
			}
			if maxCount < count {
				maxCount = count
			}
		}
	}
	return maxCount
}
