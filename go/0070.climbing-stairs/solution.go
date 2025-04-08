// Created by cubatic45 at 2025/04/08 09:41
// leetgo: 1.4.13
// https://leetcode.cn/problems/climbing-stairs/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func climbStairs(n int) (ans int) {
	if n <= 2 {
		return n
	}
	fn := make([]int, n+1)
	fn[1] = 1
	fn[2] = 2
	for i := 3; i <= n; i++ {
		fn[i] = fn[i-1] + fn[i-2]
	}
	ans = fn[n]

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	ans := climbStairs(n)

	fmt.Println("\noutput:", Serialize(ans))
}
