package leetcode_go

import "strconv"

// 解法一: 转换为字符串
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	length := len(s)
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}

// 解法2: 反转x,可参考 07.reverse integer
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	result := 0
	y := x
	for y != 0 {
		result = result*10 + y%10
		y /= 10
	}
	return result == x
}

