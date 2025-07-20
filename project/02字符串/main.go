package main

import "fmt"

////=====================================================================================================

// 字符串 有效的括号 考察：字符串处理、栈的使用 题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效 链接：https://leetcode-cn.com/problems/valid-parentheses/
func isValid(s string) bool {
	run := []rune(s)
	result := false
	length := len(run)
	if length == 1 {
		return false
	}
	a1 := "("
	run1 := []rune(a1)
	run1_val := run1[0]
	a2 := ")"
	run2 := []rune(a2)
	run2_val := run2[0]
	a3 := "{"
	run3 := []rune(a3)
	run3_val := run3[0]
	a4 := "}"
	run4 := []rune(a4)
	run4_val := run4[0]
	a5 := "["
	run5 := []rune(a5)
	run5_val := run5[0]
	a6 := "]"
	run6 := []rune(a6)
	run6_val := run6[0]
	if run[0] == run2_val || run[0] == run4_val || run[0] == run6_val {
		return false
	}
	if run[length-1] == run1_val || run[length-1] == run3_val || run[length-1] == run5_val {
		return false
	}
	for k, v := range run {
		result = false
		run1_cnt := 0
		run2_cnt := 0
		run3_cnt := 0
		run4_cnt := 0
		run5_cnt := 0
		run6_cnt := 0
		for k1, v1 := range run {
			switch v {
			case run1_val:
				if k1 > k {
					if v1 == run3_val {
						run3_cnt += 1
					} else if v1 == run4_val {
						run4_cnt += 1
					} else if v1 == run5_val {
						run5_cnt += 1
					} else if v1 == run6_val {
						run6_cnt += 1
					}
				}
				if v1 == run2_val && ((k1-k)%2 > 0 || k1 > k && k+k1 == length-1) && run3_cnt == run4_cnt && run5_cnt == run6_cnt {
					result = true
					break
				}
			case run2_val:
				if v1 == run1_val && ((k-k1)%2 > 0 || k1 < k && k+k1 == length-1) {
					result = true
					break
				}
			case run3_val:
				if k1 > k {
					if v1 == run1_val {
						run1_cnt += 1
					} else if v1 == run2_val {
						run2_cnt += 1
					} else if v1 == run5_val {
						run5_cnt += 1
					} else if v1 == run6_val {
						run6_cnt += 1
					}
				}
				if v1 == run4_val && ((k1-k)%2 > 0 || k1 > k && k+k1 == length-1) && run1_cnt == run2_cnt && run5_cnt == run6_cnt {
					result = true
					break
				}
			case run4_val:
				if v1 == run3_val && ((k-k1)%2 > 0 || k1 < k && k+k1 == length-1) {
					result = true
					break
				}
			case run5_val:
				if k1 > k {
					if v1 == run1_val {
						run1_cnt += 1
					} else if v1 == run2_val {
						run2_cnt += 1
					} else if v1 == run3_val {
						run3_cnt += 1
					} else if v1 == run4_val {
						run4_cnt += 1
					}
				}
				if v1 == run6_val && ((k1-k)%2 > 0 || k1 > k && k+k1 == length-1) && run1_cnt == run2_cnt && run3_cnt == run4_cnt {
					result = true
					break
				}
			case run6_val:
				if v1 == run5_val && ((k-k1)%2 > 0 || k1 < k && k+k1 == length-1) {
					result = true
					break
				}
			}
		}
		if !result {
			break
		}
	}
	return result

}

// 最长公共前缀 考察：字符串处理、循环嵌套 题目：查找字符串数组中的最长公共前缀 链接：https://leetcode-cn.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	result := ""
	strs_length := len(strs)
	if strs_length == 0 {
		return result
	}
	if strs_length == 1 {
		return string([]rune(strs[0]))
	}
	first_length := len([]rune(strs[0]))
outter:
	for i := 0; i < first_length; i++ {
		var each_same_run rune
		for k, v := range strs {
			run := []rune(v)
			if len(run) < i+1 {
				break outter
			}
			each_char := run[i]
			if k == 0 {
				each_same_run = each_char
			} else if k == strs_length-1 {
				if each_same_run == each_char {
					result += string(each_same_run)
				} else {
					break outter
				}
			} else {
				if each_same_run != each_char {
					break outter
				}
			}
		}
	}
	return result
}

func main() {

	// 字符串 有效的括号 考察：字符串处理、栈的使用 题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效 链接：https://leetcode-cn.com/problems/valid-parentheses/
	s := "{[]}"
	fmt.Println(isValid(s))

	//最长公共前缀 考察：字符串处理、循环嵌套 题目：查找字符串数组中的最长公共前缀 链接：https://leetcode-cn.com/problems/longest-common-prefix/
	var strs []string = []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))

}
