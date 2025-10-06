/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
*/
package main

import "fmt"

func merge(intervals [][]int) [][]int {
	// 先排序
	for i := 0; i < len(intervals)-1; i++ {
		for j := i + 1; j < len(intervals); j++ {
			if intervals[i][0] > intervals[j][0] {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}

	result := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		if len(result) == 0 {
			result = append(result, intervals[i])
		} else {
			last := result[len(result)-1]
			if intervals[i][0] <= last[1] {
				if intervals[i][1] > last[1] {
					last[1] = intervals[i][1]
				}
			} else {
				result = append(result, intervals[i])
			}
		}
	}
	return result
}

func main() {
	arr := merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	fmt.Println(arr)

	arr = merge([][]int{{1, 4}, {4, 5}})
	fmt.Println(arr)

	arr = merge([][]int{{4, 7}, {1, 4}})
	fmt.Println(arr)
}
