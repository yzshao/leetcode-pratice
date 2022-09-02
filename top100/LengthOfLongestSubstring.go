package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLongestSubstring2("dvda"))
}

func lengthOfLongestSubstring(s string) int {
	if s == "" || len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	ret := 0
	tempRet := 0
	subStr := ""
	for i := 0; i < len(s); i++ {
		if !strings.Contains(subStr, string(s[i])) {
			subStr += string(s[i])
			tempRet++
		} else {
			if ret < tempRet {
				ret = tempRet
			}
			index := strings.Index(subStr, string(s[i]))
			subStr = subStr[index+1:] + string(s[i])
			tempRet = len(subStr)
		}
	}
	if ret < tempRet {
		ret = tempRet
	}
	return ret
}

func lengthOfLongestSubstring2(s string) int {
	if s == "" || len(s) == 0 {
		return 0
	}
	recordMap := make(map[string]int)
	max := 0
	left := 0
	for i := 0; i < len(s); i++ {
		if index, ok := recordMap[string(s[i])]; ok {
			left = maxNum(left, index+1)
		}
		recordMap[string(s[i])] = i
		max = maxNum(max, i-left+1)
	}
	return max
}

func lengthOfLongestSubstring3(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
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
		ans = maxNum(ans, rk-i+1)
	}
	return ans
}

func lengthOfLongestSubstring4(s string) int {
	if len(s) == 0 {
		return 0
	}
	m := map[byte]bool{}
	n := len(s)
	l, r, ret := 0, 0, 0
	for r < n {
		for r < n && !m[s[r]] {
			// 未出现重复字符，将出现过的字符放入map中，右指针不断右移
			m[s[r]] = true
			r++
		}
		//出现重复字符或者已经遍历完所有字符
		//先对结果进行判断和赋值，当前的最大无重复字符数
		ret = maxNum(ret, r-l)
		//判断是否已经遍历完了，如果遍历完就直接退出
		if r == n {
			break
		}
		//未遍历完，那么需要处理左指针以及出现过的字符
		for l < r {
			//移除重复字符及其之前的字符
			m[s[l]] = false
			l++
			if s[l-1] == s[r] {
				break
			}
		}
	}
	return ret
}

func maxNum(x, y int) int {
	if x > y {
		return x
	}
	return y
}
