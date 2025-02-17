// Created by cubatic45 at 2025/02/13 18:15
// leetgo: 1.4.13
// https://leetcode.cn/problems/two-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func twoSum(nums []int, target int) (ans []int) {
	m := map[int]int{}
	for i, num := range nums {
		val, ok := m[target-num]
		if ok {
			ans = []int{val, i}
			return
		}
		m[num] = i
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := twoSum(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
