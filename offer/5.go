package main

import "fmt"

// 去除字符串内的空格,用%20代替

func main() {
	space := RemoveTheBlankSpace("h ell o go")
	fmt.Println(space)
}

func RemoveTheBlankSpace(s string) string {
	//获取字符串长度
	length := 0
	////获取空格数量
	blankSpaceCount := 0
	for i := range s {
		length++
		if s[i] == 32 {
			blankSpaceCount++
		}
	}

	//替换后的字符串长度
	length2 := length + blankSpaceCount*2
	var s2 = make([]byte, length2)
	point1 := length - 1
	point2 := length2 - 1

	for point1 >= 0 && point2 >= point1 {
		u := s[point1]
		if u == 32 {
			//空字符串
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
