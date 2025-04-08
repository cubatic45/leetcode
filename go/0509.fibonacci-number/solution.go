// Created by cubatic45 at 2025/04/08 11:23
// leetgo: 1.4.13
// https://leetcode.cn/problems/fibonacci-number/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// fib(n) = fib(n-1)+fib(n-2)
func fib(n int) (ans int) {
	f0, f1 := 0, 1
	for range n {
		f0, f1 = f0+f1, f0
	}
	return f0
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	ans := fib(n)

	fmt.Println("\noutput:", Serialize(ans))
}
