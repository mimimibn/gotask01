package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	func1()
	func2()
	func3()
	func4()
	func5()
	func6()
	func7()
	func8()
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

/*
*
有效的括号 ,
考察：字符串处理、栈的使用

题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
*/
func func3() {
	//	一个括号的开始只能和下一个或者是最后一个是一对才能是有效
	str := "([)]"
	dataMap := map[string]string{}
	dataMap[")"] = "("
	dataMap["]"] = "["
	dataMap["}"] = "{"
	//由于使用率byte，所以这里只支持英文括号
	bytes := []byte(str)
	num := 0
	for i := 0; i < len(bytes); i++ {
		lastIndex := (len(bytes) - 1) - i
		next := i + 1
		if next > len(bytes)-1 {
			break
		}
		if dataMap[string(bytes[i])] != "" {
			continue
		}
		if string(bytes[i]) == dataMap[string(bytes[lastIndex])] || string(bytes[i]) == dataMap[string(bytes[next])] {
		} else {
			num++
			break
		}
	}
	if num > 0 {
		fmt.Println(false)
	} else {
		fmt.Println(true)
	}
}

/*
*
考察：字符串处理、循环嵌套

题目：查找字符串数组中的最长公共前缀
*/
func func4() {
	strs := []string{"flower", "flow", "flight"}
	temp := []byte(strs[0])
	numTemp := []int{}
	for i, v := range strs {
		if i == 0 {
			continue
		}
		bytes := []byte(v)
		num := -1
		for j, v := range bytes {
			if j <= len(temp) && string(temp[j]) == string(v) {
				num++
			} else {
				break
			}
		}
		numTemp = append(numTemp, num)
	}
	min := numTemp[0]
	for _, v := range numTemp[1:] {
		if v < min {
			min = v
		}
	}
	for i, v := range temp {
		if i <= min {
			fmt.Print(string(v))
		}
	}
}

/*
*
加一

难度：简单

考察：数组操作、进位处理

题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
*/
func func5() {
	digits := []int{9, 9, 8}
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] < 10 {
			break
		}
		digits[i] = 0
	}
	if digits[0] == 0 {
		ints := []int{1}
		ints = append(ints, digits...)
		fmt.Println(ints)
	} else {
		fmt.Println(digits)
	}
}

/*
*
26. 删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
*/
func func6() {
	nums := []int{0, 1, 1, 2, 3, 4, 4}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	fmt.Println(i + 1)
}

/*
*
56. 合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，
将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
*/
func func7() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	//temp := []int{}
	//tempIndex := 0
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := [][]int{intervals[0]}
	for _, v := range intervals[1:] {
		end := merged[len(merged)-1]
		if v[0] <= end[1] {
			end[1] = v[1]
		} else {
			merged = append(merged, v)
		}
	}
	fmt.Println(merged)
}

/*
*两数之和

考察：数组遍历、map使用

题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
*/
func func8() {
	nums := []int{1, 5, 3, 4, 0}
	target := 3
	dataMap := map[int]int{}
	for _, v := range nums {
		d := target - v
		if dataMap[d] == 0 && dataMap[v] == 0 {
			dataMap[d] = v
		} else {
			fmt.Println(d, v)
		}
	}
}
