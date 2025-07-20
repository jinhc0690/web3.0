package main

import "fmt"

////=====================================================================================================

// 基础 两数之和 考察：数组遍历、map使用 题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数 链接：https://leetcode-cn.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	length := len(nums)
	result := []int{}
outter:
	for k, v := range nums {
		for i := k + 1; i < length; i++ {
			if v+nums[i] == target {
				result = append(result, k, i)
				break outter
			}
		}
	}
	return result
}

func main() {

	// 基础 两数之和 考察：数组遍历、map使用 题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数 链接：https://leetcode-cn.com/problems/two-sum/
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))

}
