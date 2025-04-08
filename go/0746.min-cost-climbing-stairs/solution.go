// Created by cubatic45 at 2025/04/08 11:33
// leetgo: 1.4.13
// https://leetcode.cn/problems/min-cost-climbing-stairs/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// f(n) = mix(f(n-1)+cost(n-1), f(n-2)+cost(n-1))
func minCostClimbingStairs(cost []int) (ans int) {
	if len(cost) == 0 {
		return 0
	}
	if len(cost) == 1 {
		return cost[0]
	}
	if len(cost) == 2 {
		return min(cost[0], cost[1])
	}
	f0, f1 := 0, 0
	for i := 2; i < len(cost)+1; i++ {
		f1, f0 = min(f0+cost[i-2], f1+cost[i-1]), f1
	}
	return f1
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	cost := Deserialize[[]int](ReadLine(stdin))
	ans := minCostClimbingStairs(cost)

	fmt.Println("\noutput:", Serialize(ans))
}
