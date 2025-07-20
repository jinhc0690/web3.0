package main

import (
	"fmt"
	"strconv"
)

////=====================================================================================================

// 136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
func singleNumber(nums []int) int {
	nummap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		_, exists := nummap[num]
		if exists {
			delete(nummap, num)
		} else {
			nummap[num] = 1
		}
	}
	var result int
	for k := range nummap {
		result = k
	}
	return result
}

// 回文数 考察：数字操作、条件判断 题目：判断一个整数是否是回文数
func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	run := []rune(str)
	lenth := len(run)
	index := lenth - 1
	cnt := lenth/2 + lenth%2
	result := true
	for k, v := range str {
		if a := run[index-k]; v != a {
			result = false
		}
		if !result {
			break
		}
		if k == cnt-1 {
			break
		}
	}
	return result
}

func main() {
	// 136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
	var nums []int = []int{2, 2, 1}
	fmt.Println(singleNumber(nums))

	// 回文数 考察：数字操作、条件判断 题目：判断一个整数是否是回文数
	x := 121
	// x := 123
	fmt.Println(isPalindrome(x))

}
