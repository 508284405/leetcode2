package main

import (
	"fmt"
	"sort"
)

/**
给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
*/
func main() {
	//sum := threeSum([]int{-1, 1, 0, -1, -3, 2})
	//sum := threeSum([]int{-3, -2, -2, -1, -1, 0, 1, 2})
	sum := threeSum([]int{0,0,0,0})
	fmt.Print(sum)
}

func threeSum(nums []int) [][]int {
	//1.0 排序
	sort.Ints(nums)
	//2.0 循环
	var res = make([][]int, 0)

	for i := range nums {
		i2 := i - 1
		if i2 <=0 {
			i2 = 0
		}
		if nums[i] > 0 && nums[i2] == nums[i] {
			continue
		}
		//定义左指针
		l := i + 1
		r := len(nums) - 1
		for l < r {
			i2 := nums[i] + nums[l] + nums[r]
			if i2 == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				break
			} else if i2 < 0 {
				//小了,做指针右移
				l++
			} else {
				//大了,右指针左移
				r--
			}
		}
	}
	return res
}
