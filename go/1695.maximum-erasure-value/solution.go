// Created by cubatic45 at 2025/07/22 11:58
// leetgo: 1.4.14
// https://leetcode.cn/problems/maximum-erasure-value/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maximumUniqueSubarray(nums []int) (ans int) {
	if len(nums) == 0 {
		return
	}
	existsMap := map[int]int{}

	sum := 0
	l := 0
	for i, v := range nums {
		r, ok := existsMap[v]
		for ; l <= r && ok; l++ {
			sum -= nums[l]
		}
		sum += v
		existsMap[v] = i
		ans = max(ans, sum)
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := maximumUniqueSubarray(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
