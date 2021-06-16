package main

import "fmt"

/**
    二维数组查找
	在一个二维数组中,每一行都按照从左到右递增的顺序排序,每一列都按照从上到下递增的顺序排序.
	请完成一个函数,输入这样的一个二维数组和一个整数,判断数组中是否含有该整数
*/

func main() {
	var array [][]int
	array = append(array, []int{1, 2, 8, 9})
	array = append(array, []int{2, 4, 9, 12})
	array = append(array, []int{4, 7, 10, 13})
	array = append(array, []int{6, 8, 11, 15})
	dimensionalArray := findNumberIn2DArray(array, 10)
	fmt.Print(dimensionalArray)
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	rows := len(matrix)
	columns := len(matrix[0])
	if matrix == nil || rows == 0 || columns == 0 {
		return false
	}
	row := 0
	column := columns - 1

	for row < rows && column > 0 {
		//1.0 取右上角的元素
		a := matrix[row][column]
		if a == target {
			return true
		} else if a > target {
			//向左移动一格
			column--
		} else {
			//向下移动一格
			row++
		}
	}
	return false
}
