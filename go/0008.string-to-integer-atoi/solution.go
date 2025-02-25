// Created by cubatic45 at 2025/02/24 17:33
// leetgo: 1.4.13
// https://leetcode.cn/problems/string-to-integer-atoi/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	abs, sign, i, n := 0, 1, 0, len(s)
	// 标记正负号
	if i < n {
		switch s[i] {
		case '-':
			sign = -1
			i++
		case '+':
			sign = 1
			i++
		}
	}
	for i < n && s[i] >= '0' && s[i] <= '9' {
		abs = 10*abs + int(s[i]-'0')
		if sign*abs < math.MinInt32 {
			return math.MinInt32
		} else if sign*abs > math.MaxInt32 {
			return math.MaxInt32
		}
		i++
	}
	return sign * abs
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := myAtoi(s)

	fmt.Println("\noutput:", Serialize(ans))
}
