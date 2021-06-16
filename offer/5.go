package main

import "fmt"

/**
去除字符串内的空格
核心: 空间替换时间
*/

func main() {
	space := ReplaceBlankSpace("hello word")
	fmt.Println(space)
}

func ReplaceBlankSpace(s string) string {
	//计算字符串的长度和字符串内的空格数量
	l := 0
	b := 0
	for i := range s {
		l++
		if s[i] == 32 {
			b++
		}
	}

	//定义替换后的字符串长度,1个空格替换成%20
	l2 := l + b*2
	var s2 = make([]byte, l2)

	point1 := l - 1
	point2 := l2 - 1
	for point1 >= 0 && point2 >= point1 {
		if s[point1] == ' ' {
			s2[point2] = '0'
			point2--
			s2[point2] = '2'
			point2--
			s2[point2] = '%'
			point2--
		} else {
			s2[point2] = s[point1]
			point2--
		}
		point1--
	}
	return string(s2)
}
