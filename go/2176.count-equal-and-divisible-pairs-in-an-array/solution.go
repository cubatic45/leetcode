// Created by cubatic45 at 2025/04/17 16:10
// leetgo: 1.4.13
// https://leetcode.cn/problems/count-equal-and-divisible-pairs-in-an-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func countPairs(nums []int, k int) (ans int) {
	if len(nums) < 2 || k == 0 {
		return
	}
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] != nums[j] {
				continue
			}
			if (i*j)%k == 0 {
				ans++
			}
		}
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := countPairs(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
