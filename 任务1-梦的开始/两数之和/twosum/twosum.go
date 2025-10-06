/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
你可以按任意顺序返回答案。
*/
package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := numsMap[target-nums[i]]; ok {
			return []int{v, i}
		}
		numsMap[nums[i]] = i
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
