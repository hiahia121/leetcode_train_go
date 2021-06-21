package main

import "fmt"

// 题目内容：有效的数组异位词； 详细内容：https://mp.weixin.qq.com/s/vM6OszkM6L1Mx2Ralm9Dig

func availdTwoStr(str1 string, str2 string) bool{
	nums1 := [26]int{}
	for _, val := range str1{
		index := val - 'a'
		nums1[index]++
	}
	for _, val := range str2{
		index := val - 'a'
		nums1[index]--
	}

	for _, val := range nums1{
		if val != 0{
			return false
		}
	}
	return true
}


func main(){
	//s1 := "anagram"
	//s2 := "nagaram"

	s1 := "rat"
	s2 := "car"
	fmt.Println(availdTwoStr(s1, s2))
}