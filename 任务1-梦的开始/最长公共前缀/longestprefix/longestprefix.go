/*
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
*/
package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	prefix := strs[0]
	for _, str := range strs {
		rstr := []rune(str)
		temp := []rune{}
		for pi, ps := range prefix {
			if pi >= len(rstr) || ps != rstr[pi] {
				break
			}
			temp = append(temp, ps)
		}
		prefix = string(temp)
	}
	return prefix
}

func main() {

}
