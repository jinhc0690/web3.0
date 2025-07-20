package main

import "fmt"

////=====================================================================================================

// 引用类型：切片 26. 删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	nummap := map[int]int{}
	result := []int{}
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		_, exists := nummap[num]
		if !exists {
			nummap[num] = 1
			result = append(result, num)
		}
	}
	for k, v := range result {
		nums[k] = v
	}
	return len(result)
}

// 56. 合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
func merge(intervals [][]int) [][]int {

	length := len(intervals)
	if length == 1 {
		return intervals
	}
	result := [][]int{}
	for k1, v1 := range intervals {
		start_number1 := v1[0]
		for k2, v2 := range intervals {
			start_number2 := v2[0]
			if start_number1 < start_number2 && k1 > k2 {
				copy := v1
				intervals[k1] = v2
				v1 = intervals[k1]
				intervals[k2] = copy
			}
		}
	}
	start_number := 0
	end_number := 0
	for k, v := range intervals {
		start_number_each := v[0]
		end_number_each := v[1]
		if k == 0 {
			start_number = start_number_each
			end_number = end_number_each
		} else {
			if start_number_each >= start_number && start_number_each <= end_number {
				if end_number_each > end_number {
					end_number = end_number_each
				}
				if k == length-1 {
					result_each := []int{start_number, end_number}
					result = append(result, result_each)
				}
			} else {
				result_each := []int{start_number, end_number}
				result = append(result, result_each)
				start_number = start_number_each
				end_number = end_number_each
				if k == length-1 {
					result_each := []int{start_number, end_number}
					result = append(result, result_each)
				}
			}
		}
	}
	return result
}

func main() {

	// 引用类型：切片 26. 删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))

	// 56. 合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
	intervals := [][]int{{4, 5}, {1, 4}, {0, 1}}
	fmt.Println(merge(intervals))

}
