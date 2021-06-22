package main

/**
输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
前序遍历 preorder = [3,9,20,15,7]  先访问子节点
中序遍历 inorder = [9,3,15,20,7]	 中间访问子节点
-- 后续遍历 [9,15,7,20,3]          最后访问子节点
	3
   / \
  9  20
	/  \
   15   7

0 <= 节点个数 <= 5000
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return nil
}
