package main

import "fmt"

// 题目：两数之和，详细内容：https://leetcode-cn.com/problems/two-sum/submissions/
func twoNumsSum(nums []int, target int) []int {
	tmp_map := make(map[int]int)
	for index2, val := range nums{
		if _, ok := tmp_map[target - val]; ok{
			index1 := tmp_map[target - val]
			return []int{index1, index2}
		}
		tmp_map[val] = index2
	}
	return []int{}
}

func main()  {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoNumsSum(nums, target))
}