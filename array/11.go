package main

import (
	"fmt"
)

/**
核心:  双指针法: 从数组的两边作为指针及边界,移动较小的边界的指针(因为边界大的指针,只能向左移动,值只会越来越小)
*/
func main() {
	area := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Print(area)
}

func maxArea(height []int) int {
	//定义左右指针,作为容器的边界
	l := 0
	r := len(height) - 1
	var c = 0
	for r > l {
		//计算容器值
		li := height[l]
		ri := height[r]
		area := 0
		if li <= ri {
			area = li * (r - l)
		} else {
			area = ri * (r - l)
		}
		if c <= area {
			c = area
		}
		//比较左右边界的大小,移动较小边界的指针
		if li <= ri {
			l++
		} else {
			r--
		}

	}
	return c
}
