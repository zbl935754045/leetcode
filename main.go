package main

import (
	"fmt"
	"reflect"
)

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30}}
	fmt.Println(matrix[4])
	//ok := findNumberIn2DArray(matrix, 5)
	//fmt.Print(ok)

	fmt.Println(reflect.TypeOf('e'))
}

func replaceSpace(s string) string {
	str := ""
	for _, sVar := range s{
		if sVar == ' ' {
			str += "%20"
		} else {
			str += string(sVar)
		}
	}
	return str
}


func findNumberIn2DArray2(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	//从右上角开始
	lenY := len(matrix)
	Y := 0
	lenX := len(matrix[0])
	X := lenX - 1
	for Y < lenY && X >= 0 {
		if matrix[Y][X] == target {
			return true
		} else if matrix[Y][X] < target {
			Y ++
		} else {
			X --
		}
	}
	return false
}


//暴力解法，for循环
func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	numMap := map[int]int{}
	for i, num := range nums {
		if value, ok := numMap[target-num]; ok {
			return []int{i, value}
		}
		numMap[num] = i
	}
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		sum := n1 + n2 + carry
		sum = sum % 10
		carry = sum / 10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}
func lengthOfLongestSubstring(s string) int {
	//记录字符是否存在过
	m := map[byte]int{}
	//记录字符串的长度
	n := len(s)
	//右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func reverse(x int) int {
	var res int
	for x != 0 {
		if temp := int32(res); (temp*10)/10 != temp {
			return 0
		}
		res = res*10 + x%10
		x = x / 10
	}
	return res
}

//剑指 Offer 04. 二维数组中的查找
func findNumberIn2DArray(matrix [][]int, target int) bool {
	for _, matrixOne := range matrix {
		for _, num := range matrixOne {
			fmt.Println(num)
			if num == target {
				return true
			}
		}
	}
	return false
}

