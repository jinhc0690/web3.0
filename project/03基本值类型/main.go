package main

import "fmt"

////=====================================================================================================

// 基本值类型 加一 难度：简单 考察：数组操作、进位处理 题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一 链接：https://leetcode-cn.com/problems/plus-one/
func plusOne(digits []int) []int {
	length := len(digits)
	extra_add := 0
	for i := 0; i < length; i++ {
		sum := 0
		index := length - 1 - i
		value := digits[index]
		if i == 0 {
			sum = value + 1
		} else {
			sum = value + extra_add
		}
		if sum < 10 {
			digits[index] = sum
			break
		}
		digits[index] = 0
		extra_add = 1
		if i == length-1 {
			extra_add = 2
		}
	}
	if extra_add == 2 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {

	// 基本值类型 加一 难度：简单 考察：数组操作、进位处理 题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一 链接：https://leetcode-cn.com/problems/plus-one/
	digits := []int{4, 3, 2, 1}
	fmt.Println(plusOne(digits))

}
