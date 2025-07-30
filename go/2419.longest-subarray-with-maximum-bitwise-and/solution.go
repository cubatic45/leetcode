// Created by cubatic45 at 2025/07/30 15:18
// leetgo: 1.4.14
// https://leetcode.cn/problems/longest-subarray-with-maximum-bitwise-and/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func longestSubarray(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	maxLen, maxNum := 0, 0
	i := 0
	for i < n {
		if nums[i] < maxNum {
			i++
			continue
		} else if nums[i] > maxNum {
			maxNum = nums[i]
			maxLen = 1
		}
		j := i
		tmpMaxLen := 0
		for j < len(nums) {
			if nums[j] == nums[i] {
				j++
				tmpMaxLen++
			} else {
				break
			}
		}
		i = j
		maxLen = max(maxLen, tmpMaxLen)
	}
	return maxLen
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := longestSubarray(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
