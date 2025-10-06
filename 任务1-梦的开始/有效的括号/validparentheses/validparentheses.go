/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：
1. 左括号必须用相同类型的右括号闭合。
2. 左括号必须以正确的顺序闭合。
3. 每个右括号都有一个对应的相同类型的左括号。
*/
package main

import "fmt"

type runeStack struct {
	data []rune
}

func (rs *runeStack) push(c rune) {
	rs.data = append(rs.data, c)
}

func (rs *runeStack) pop() rune {
	top := rs.data[len(rs.data)-1]
	rs.data = rs.data[:len(rs.data)-1]
	return top
}

func (rs *runeStack) isEmpty() bool {
	return len(rs.data) == 0
}

// 只保留题目中的特殊字符，然后使用
func isValid(s string) bool {
	charMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := runeStack{}

	for _, v := range s {
		if _, ok := charMap[v]; ok {
			stack.push(v)
		} else if v == ')' || v == ']' || v == '}' {
			if stack.isEmpty() {
				return false
			}
			if charMap[stack.pop()] != v {
				return false
			}
		}
	}

	return stack.isEmpty()
}

func main() {
	s := "()"
	fmt.Println(isValid(s))
}
