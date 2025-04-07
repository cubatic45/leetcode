// Created by cubatic45 at 2025/02/26 11:01
// leetgo: 1.4.13
// https://leetcode.cn/problems/palindrome-number/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	xStr := strconv.Itoa(x)
	for i, j := 0, len(xStr)-1; i < j; {
		if xStr[i] != xStr[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	x := Deserialize[int](ReadLine(stdin))
	ans := isPalindrome(x)

	fmt.Println("\noutput:", Serialize(ans))
}
