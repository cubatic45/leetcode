// Created by cubatic45 at 2025/02/17 10:15
// leetgo: 1.4.13
// https://leetcode.cn/problems/longest-palindromic-substring/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
/*
bbbb
*/

func longestPalindrome(s string) string {
	ans := ""
	for i := 0; i < len(s); i++ {
		for j := i; j <= i+1 && j < len(s); j++ {
			tmp := 0
			for i-tmp >= 0 && j+tmp < len(s) {
				if s[i-tmp] == s[j+tmp] {
					if tmp*2+j-i+1 > len(ans) {
						ans = s[i-tmp : j+tmp+1]
					}
					tmp++
				} else {
					break
				}
			}
		}
	}
	return ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := longestPalindrome(s)
	fmt.Println("\noutput:", Serialize(ans))
}
