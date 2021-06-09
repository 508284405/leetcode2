package main

import (
	"fmt"
	"math"
)

//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

func main() {
	median := getMedian([]int{1, 2, 3, 4}, []int{2, 5, 6, 7, 8, 9})
	fmt.Println(median)
}

//二分法求中位数:  二分法分的是小数组
//假设合并数组以后，中位数左边的个数 int totalLeft = (m + n + 1)/2  左边永远比右边多一个
//1 2 3 4
//2 5 6 7 8 9  中位数是4.5
func getMedian(nums1 []int, nums2 []int) float64 {
	//1.0 让nums1成为小的数组
	if len(nums1) > len(nums2) {
		var temp []int = nums2
		nums2 = nums1
		nums1 = temp
	}
	//2.0 取出nums1和nums2 的长度
	m := len(nums1)
	n := len(nums2)

	//3.0 计算合并后的分组的左半部分长度
	totalLeft := (m + n + 1) / 2

	//从小的数组开始划分
	//nums1的划分范围
	l := 0
	r := m

	//保证nums1的分割线左边元素小于等于nums2的分割线右边元素 && num2的分割线左边元素小于等于num1的分割线右边元素
	for l < r {
		//nums1的分割线
		i := l + (r-l+1)/2
		//nums2的分割线
		j := totalLeft - i
		//条件取反
		//nums1的分割线左边元素大于nums2的分割线右边元素== nums1的分割线需要向左移动
		if nums1[i-1] > nums2[j] {
			r = i - 1
		} else {
			//搜索区间在右半部分
			l = i
		}
	}
	//找到分割线以后,l = r = 分割线,l 表示第一个数组的小元素
	var i = r
	//找到nums2的分割线
	var j = totalLeft - i

	var nums1LeftMax int
	if i == 0 {
		//nums1分割线左边没有元素
		nums1LeftMax = INT_MAX
	} else {
		nums1LeftMax = nums1[i-1]
	}

	var nums1RightMin int
	if i == m {
		//nums1右边没有元素
		nums1RightMin = INT_MIN
	} else {
		nums1RightMin = nums1[i]
	}

	var nums2LeftMax int
	if j == 0 {
		//nums2分割线左边没有元素
		nums2LeftMax = INT_MIN
	} else {
		nums2LeftMax = nums2[j-1]
	}

	var nums2RightMin int

	if j == n {
		nums2RightMin = INT_MAX
	} else {
		nums2RightMin = nums2[j]
	}
	if (m+n)%2 == 1 {
		// 如果两个数组的长度之和为奇数，直接返回两个数组在分割线左边的最大值即可
		return math.Max(float64(nums1LeftMax), float64(nums2LeftMax))
	} else {
		// 如果两个数组的长度之和为偶数，返回的是两个数组在左边的最大值和两个数组在右边的最小值的和的二分之一
		// 此处不能被向下取整，所以要强制转换为double类型
		return (math.Max(float64(nums1LeftMax), float64(nums2LeftMax)) + math.Min(float64(nums1RightMin), float64(nums2RightMin))) / 2
	}
}

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX
