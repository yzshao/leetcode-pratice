package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
输入：s = "3[a]2[bc]"
输出："aaabcbc"
*/
func main() {
	fmt.Println(decodeString2("3[a2[bc]]"))
}

var (
	src string
	p   int
)

func decodeString2(s string) string {
	src = s
	p = 0
	return getString2()
}

func getString2() string {
	// 递归终止条件 1、src所有字符遍历完成  2、p指针指向字符']'
	if p == len(src) || src[p] == ']' {
		return ""
	}
	repeatTime := 1
	cur := src[p]
	ret := ""
	if cur >= '0' && cur <= '9' {
		repeatTimeStr := getDigits(src, &p)
		p++
		s := getString2()
		p++
		rTime, _ := strconv.Atoi(repeatTimeStr)
		repeatTime = rTime
		ret = strings.Repeat(s, repeatTime)
	} else if (cur >= 'a' && cur <= 'z') || (cur >= 'A' && cur <= 'Z') {
		ret += string(cur)
		p++
	}
	return ret + getString2()
}

func decodeString(s string) string {
	resStack := make([]string, 0)
	var ptr = 0
	for ptr < len(s) {
		cur := s[ptr]
		if cur >= '0' && cur <= '9' {
			// 数字类型
			digits := getDigits(s, &ptr)
			resStack = append(resStack, digits)
		} else if (cur >= 'a' && cur <= 'z') || (cur >= 'A' && cur <= 'Z') || cur == '[' {
			resStack = append(resStack, string(cur))
			ptr++
		} else {
			// ']' 情况
			strSlice := make([]string, 0)
			index := len(resStack) - 1
			// 1、从stack中弹出字符
			for ; resStack[index] != "[" && index >= 0; index-- {
				strSlice = append(strSlice, resStack[index])
			}
			// 跳过'['
			index--
			// 将重复次数拿出来
			str := reverseString(strSlice)
			repeatTime, _ := strconv.Atoi(resStack[index])
			t := strings.Repeat(str, repeatTime)
			if index == 0 {
				resStack = resStack[:0]
			} else {
				resStack = resStack[0:index]
			}
			// 2、将str压入stack中
			resStack = append(resStack, t)
			ptr++
		}
	}
	return getString(resStack)
}

func getDigits(s string, ptr *int) string {
	ret := ""
	for {
		if s[*ptr] <= '9' && s[*ptr] >= '0' {
			ret += string(s[*ptr])
			*ptr++
		} else {
			break
		}
	}
	return ret
}

func reverseString(s []string) string {
	runes := s
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return getString(runes)
}

func getString(sList []string) string {
	ret := ""
	for _, s := range sList {
		ret += s
	}
	return ret
}
