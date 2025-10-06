// 任务1-梦的开始 题目1：只出现一次的数字
package main

import "fmt"

// 查找只出现一次的数组
func singleNumber(nums []int) int {
	numsMap := make(map[int]int)
	for _, v := range nums {
		numsMap[v]++
	}

	for k, v := range numsMap {
		if v == 1 {
			return k
		}
	}

	return -1
}

func main() {
	nums := []int{1}
	s := singleNumber(nums)
	fmt.Println(s)
}
