package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

// 只出现一次的数字
func singleNumber(nums []int) int {
	var result int
	for i := 0; i < len(nums); i++ {
		countnum := 0
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] {
				countnum++
			}
		}
		if countnum == 1 {
			result = nums[i]
		}
	}
	return result
}

// 回文数
func isPalindrome(x int) bool {
	str1 := strconv.Itoa(x)
	str2 := ReverseString(str1)
	return str1 == str2
}
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 有效的括号
func isValid(s string) bool {
	slice := []rune(s)
	if len(slice)%2 != 0 {
		return false
	} else {
		for i := 0; i < len(slice)-1; i++ {
			switch string(slice[i]) {
			case "(":
				if string(slice[i+1]) == ")" {
					slice = slices.Delete(slice, i, i+2)
					i = -1
					continue
				}
			case "[":
				if string(slice[i+1]) == "]" {
					slice = slices.Delete(slice, i, i+2)
					i = -1
					continue
				}
			case "{":
				if string(slice[i+1]) == "}" {
					slice = slices.Delete(slice, i, i+2)
					i = -1
					continue
				}
			}
		}
		if len(slice) == 0 {
			return true
		} else {
			return false
		}
	}
}

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	strLenth := []int{}
	strcommom := []string{}
	count := 0
	for i := 0; i < len(strs); i++ {
		strLenth = append(strLenth, len(strs[i]))
	}
	exitFlag := false
	for i := 0; i < slices.Min(strLenth); i++ {
		for j := 0; j < len(strLenth)-1; j++ {
			if strs[j][i] == strs[j+1][i] {
				count = j
				continue
			} else {
				if i == 0 {
					return ""
				} else {
					exitFlag = true
					break
				}
			}
		}
		if exitFlag {
			break
		}
		if count == len(strLenth)-2 {
			strcommom = append(strcommom, string(strs[count][i]))
		}
	}
	result := strings.Join(strcommom, "")
	return result
}

// 加一
func plusOne(digits []int) []int {
	suminit := 0
	digitplus := []int{}
	for i := 0; i < len(digits); i++ {
		suminit += digits[i] * int(math.Pow(10, float64(len(digits)-i-1)))
	}
	suminit += 1
	suminitLen := len(strconv.Itoa(suminit))
	for i := 0; i < suminitLen; i++ {
		cf := math.Pow(10, float64(suminitLen-i-1))
		num := suminit / int(cf)
		suminit -= num * int(cf)
		digitplus = append(digitplus, num)
	}
	return digitplus
}

// 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	slice := nums[:]
	newnums := []int{}
	newnums = append(newnums, nums[0])
	for i, j := 1, 0; i < len(slice) && j < len(slice); j++ {
		if slice[i] == slice[j] {
			slice = slices.Delete(slice, j, j+1)
			j--
		} else {
			newnums = append(newnums, slice[i])
			i++
		}
	}
	copy(nums, newnums)
	return len(newnums)
}

// 两数之和
func twoSum(nums []int, target int) []int {
	lenNum := len(nums)
	for i := 0; i < lenNum; i++ {
		for j := i + 1; j < lenNum; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 合并区间
func merge(intervals [][]int) [][]int {
	slice := make([][]int, len(intervals))
	for i := range intervals {
		slice[i] = intervals[i][:]
	}
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			fmt.Println()
			fmt.Print("当前比较的两个区间是第", i, "个区间和第", j, "个区间：[", slice[i][0], slice[i][1], "]", "[", slice[j][0], slice[j][1], "]")
			if max(slice[i][0], slice[j][0]) <= min(slice[i][1], slice[j][1]) {
				fmt.Println("，区间相交，可以合并")
				newinterval := make([]int, 2)
				if slice[i][0] >= slice[j][0] {
					newinterval[0] = slice[j][0]
				} else {
					newinterval[0] = slice[i][0]
				}
				// fmt.Println(newinterval[0])
				if slice[i][1] >= slice[j][1] {
					newinterval[1] = slice[i][1]
				} else {
					newinterval[1] = slice[j][1]
				}
				fmt.Print("合并后的区间是：", newinterval)
				slice[0] = newinterval
				slice = slices.Delete(slice, j, j+1)
				fmt.Println("合并后的二维切片是：", slice)
				j--
			} else {
				fmt.Println("区间相离，继续比较")
				continue
			}

		}
	}
	return slice
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}
