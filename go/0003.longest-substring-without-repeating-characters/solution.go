// Created by cubatic45 at 2025/02/14 10:14
// leetgo: 1.4.13
// https://leetcode.cn/problems/longest-substring-without-repeating-characters/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func lengthOfLongestSubstring(s string) (ans int) {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}
	i, j := 0, 1
	m := map[byte]int{s[0]: 1}
	for {
		if j >= len(s) {
			ans = max(j-i, ans)
			break
		}
		v := m[s[j]]
		if v > 0 {
			ans = max(j-i, ans)
			m[s[i]]--
			i++
		} else {
			m[s[j]]++
			j++
		}
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := lengthOfLongestSubstring(s)

	fmt.Println("\noutput:", Serialize(ans))
}
