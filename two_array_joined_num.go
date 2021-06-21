package main

import "fmt"

// 题目内容, 两个数组集合求交集，具体内容见链接：https://mp.weixin.qq.com/s/N9iqAchXreSVW7zXUS4BVA

func twoArrayJoinRet(nums1 []int, nums2 []int) ([]int){
	ret := []int{}
	tmp_map := make(map[int]int)
	for _, val := range nums2{
		tmp_map[val] = 0
	}

	for _, val := range nums1{
		if _, ok := tmp_map[val]; ok{
			ret = append(ret, val)
		}
	}
	return ret
}

func main()  {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(twoArrayJoinRet(nums1, nums2))
}