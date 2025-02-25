// Created by cubatic45 at 2025/02/20 15:40
// leetgo: 1.4.13
// https://leetcode.cn/problems/zigzag-conversion/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func convert(s string, numRows int) string {
	if len(s) <= 2 || numRows == 1 {
		return s
	}
	tmp := make([][]byte, numRows)
	gap := numRows*2 - 2
	for i := 0; ; i++ {
		if i*(numRows*2-2) >= len(s) {
			break
		}
		subStr := s[i*gap : min((i+1)*gap, len(s))]
		for i := range subStr {
			tmpIndex := i % numRows
			if i/numRows > 0 {
				tmpIndex = numRows - tmpIndex - 2
			}
			tmp[tmpIndex] = append(tmp[tmpIndex], subStr[i])
		}
	}
	ans := ""
	for _, v := range tmp {
		ans += string(v)
	}

	return ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	numRows := Deserialize[int](ReadLine(stdin))
	ans := convert(s, numRows)

	fmt.Println("\noutput:", Serialize(ans))
}
