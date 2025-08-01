// Created by cubatic45 at 2025/08/01 10:03
// leetgo: 1.4.14
// https://leetcode.cn/problems/pascals-triangle/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func generate(numRows int) [][]int {
	if numRows <= 0 {
		return nil
	}
	ans := make([][]int, numRows)
	for i := range numRows {
		ans[i] = generateOne(i)
	}
	return ans
}

func generateOne(n int) []int {
	row := make([]int, n+1)
	row[0] = 1
	for i := 1; i <= n; i++ {
		// Use the relation: C(n, i) = C(n, i-1) * (n - i + 1) / i
		row[i] = row[i-1] * (n - i + 1) / i
	}
	return row
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	numRows := Deserialize[int](ReadLine(stdin))
	ans := generate(numRows)

	fmt.Println("\noutput:", Serialize(ans))
}
