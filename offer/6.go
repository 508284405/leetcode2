package main

import (
	"fmt"
)

//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
var a []int

func main() {
	var c = ListNode{3, nil}

	var b = ListNode{2, &c}

	var a = ListNode{
		key:      1,
		nextNode: &b,
	}
	ints := reversePrintStack(&a)
	fmt.Println(ints)
}

type ListNode struct {
	key      int
	nextNode *ListNode
}

//递归实现
func reversePrint(head *ListNode) []int {
	if head.nextNode != nil {
		ints := reversePrint(head.nextNode)
		return append(ints, head.key)
	} else {
		return []int{head.key}
	}
}

//栈实现
func reversePrintStack(head *ListNode) []int {
	var nums []int
	for head.nextNode != nil {
		nums = append(nums, head.key)
		head = head.nextNode
	}
	nums = append(nums, head.key)
	//反向输出数组
	var reserveNums = make([]int, len(nums))

	for i, num := range nums {
		reserveNums[len(nums)-1-i] = num
	}

	return reserveNums
}
