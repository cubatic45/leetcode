// Created by cubatic45 at 2025/02/20 16:39
// leetgo: 1.4.13
// https://leetcode.cn/problems/reverse-integer/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverse(x int) (ans int) {
	for x != 0 {
		seed := x % 10
		x = x / 10
		ans = 10*ans + seed
	}
	if ans > math.MaxInt32 || ans < math.MinInt32 {
		ans = 0
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	x := Deserialize[int](ReadLine(stdin))
	ans := reverse(x)

	fmt.Println("\noutput:", Serialize(ans))
}
