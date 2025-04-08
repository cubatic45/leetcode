// Created by cubatic45 at 2025/04/08 11:26
// leetgo: 1.4.13
// https://leetcode.cn/problems/n-th-tribonacci-number/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2
func tribonacci(n int) (ans int) {
	t0, t1, t2 := 0, 1, 1
	for range n {
		t0, t1, t2 = t1, t2, t0+t1+t2
	}
	return t0
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	ans := tribonacci(n)

	fmt.Println("\noutput:", Serialize(ans))
}
