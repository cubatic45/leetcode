// Created by cubatic45 at 2025/04/08 15:33
// leetgo: 1.4.13
// https://leetcode.cn/problems/delete-and-earn/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func deleteAndEarn(nums []int) (ans int) {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	}
	// f(n)表示删除 num = n 获得的最大点数
	// f(n) = max(f(n-1), f(n-2)+n)
	maxNum := 0
	for _, v := range nums {
		maxNum = max(maxNum, v)
	}
	s := make([]int, maxNum)
	for _, v := range nums {
		s[v-1] += v
	}
	f0, f1 := 0, 0
	for _, v := range s {
		f0, f1 = f1, max(f0+v, f1)
	}
	return max(f0, f1)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := deleteAndEarn(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
