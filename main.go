package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	func2()
}

/*
*
题目：136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

	可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*/
func func1() {
	nums := []int{4, 1, 2, 1, 2}
	temp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		temp[nums[i]] += 1
	}
	for k, v := range temp {
		if v == 1 {
			fmt.Println(k)
		}
	}
}

/*
*
回文数

考察：数字操作、条件判断
题目：判断一个整数是否是回文数
*/
func func2() {
	str := 1234321
	tempStr := strconv.Itoa(str)
	bytes := []byte(tempStr)
	var newStr strings.Builder
	for i := len(bytes) - 1; i >= 0; i-- {
		newStr.WriteString(string(bytes[i]))
	}
	if newStr.String() == tempStr {
		fmt.Println(str, "是回文数")
	} else {
		fmt.Println(str, "不是回文数")
	}
}
