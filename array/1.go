package main

import (
	"fmt"
)

//两数之和
//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//
//你可以按任意顺序返回答案。

func main() {
	hash := twoNumsSumHash([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 10)
	fmt.Println(hash)
}

func twoNumsSumHash(nums []int, target int) []int {
	//key是值 value是索引
	mapHash := make(map[int]int)
	for i, m := range nums {
		if _, ok := mapHash[target-m]; ok {
			return []int{i, mapHash[target-m]}
		}
		mapHash[m] = i
	}
	return nil
}
