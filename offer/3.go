package main

import "fmt"

//在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
//数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字

func main() {
	array := getIntFromArrayNoAlter([]int{1, 0, 2, 4, 3, 5})
	fmt.Println(array)
}

/**
	0.0  判断索引值和数字值是否一样,一样则越过
	1.0  当同时出现某个索引处的值和索引不一样,且该值作为索引,其值的值和其值得索引对的上,则该值是重复的值
	2.0  遍历数组,将数字挪到对应的索引处(前提数字和索引对不上)
    条件: 1.索引和其值不同 2. 其值作为索引,其索引和值相同
*/
func getIntFromArray(nums []int) int {
	i := 0
	for i < len(nums) {
		//创造条件1,将第一次出现的数字,挪到其对应的索引下
		if i == nums[i] {
			i++
			continue
		}
		//判断条件2
		if nums[nums[i]] == nums[i] {
			return nums[i]
		}
		//将数字挪到其对应索引的位置
		var j = nums[i]
		nums[i] = nums[j]
		nums[j] = j
	}
	return -1
}

/**
  	不修改数组
	在长度为n+1的数组内,所有数字都在1～n的范围内,请找出数组内任意一个重复的数字,但不能修改数组
*/
func getIntFromArrayNoAlter(nums []int) int {
	start := 1
	end := len(nums) - 1
	for start <= end {
		middle := (end-start)>>1 + start
		count := getRangeCount(nums, start, middle)
		if start == end {
			if count > 1 {
				return start
			} else {
				break
			}
		}

		if count > middle-start+1 {
			end = middle
		} else {
			//middle 已经计算过了,需要从下一位开始计算
			start = middle + 1
		}
	}
	return -1
}

/**
  计算数组内出现在range start and end 内的数字个数
*/
func getRangeCount(nums []int, start int, end int) int {
	var count = 0
	for num := range nums {
		if num <= end && num >= start {
			count++
		}
	}
	return count
}
