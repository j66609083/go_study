package main

import (
	"fmt"
	"strconv"
)

// 判断一个整数是否是回文数。使用转字符串的方式实现
func isPalindrome(x int) bool {
	xs := []rune(strconv.Itoa(x))
	for i, j := 0, len(xs)-1; i < j; i, j = i+1, j-1 {
		if xs[i] != xs[j] {
			return false
		}
	}
	return true
}

// 判断一个整数是否是回文数。不使用转字符串的方式实现
func isPalindromeEnhance(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	revr := 0
	for x > revr {
		revr = revr*10 + x%10
		x /= 10
	}

	return x == revr || x == revr/10
}

func main() {
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindromeEnhance(-121))
}
