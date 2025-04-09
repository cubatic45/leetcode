// Created by cubatic45 at 2025/04/08 15:19
// leetgo: 1.4.13
// https://leetcode.cn/problems/house-robber/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func rob(nums []int) (ans int) {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return max(nums[0], nums[1])
	}
	// f(n) = max(f(n-2)+nums[n], f(n-1))
	f1, f2 := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		f1, f2 = f2, max(f1+nums[i], f2)
	}

	return f2
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := rob(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
