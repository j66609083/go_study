/*
给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
将大整数加 1，并返回结果的数字数组。
https://leetcode.cn/problems/plus-one/description/
*/
package main

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}

	result := make([]int, len(digits)+1)

	needAdd := true
	for i := len(digits) - 1; i >= 0; i-- {
		if needAdd {
			if digits[i] < 9 {
				result[i+1] = digits[i] + 1
				needAdd = false
			} else {
				result[i] = 1
				result[i+1] = 0
			}
		} else {
			result[i+1] = digits[i]
		}
	}

	if result[0] == 0 {
		result = result[1:]
	}

	return result
}
