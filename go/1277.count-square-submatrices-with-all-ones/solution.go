// Created by cubatic45 at 2025/08/20 14:50
// leetgo: 1.4.14
// https://leetcode.cn/problems/count-square-submatrices-with-all-ones/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func countSquares(matrix [][]int) (ans int) {
	// 遍历，找上方和右方就不会重复
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] != 1 {
				continue
			}
			ans++
			maxExpand := min(i, len(matrix[i])-1-j)
			if maxExpand == 0 {
				continue
			}
			// 判断 matrix[i][j] 向上和向右扩展
			for expand := 1; expand <= maxExpand; expand++ {
				add := true
				// 判断 x，y 方向
				for v := 0; v <= expand; v++ {
					// 顶边
					if matrix[i-expand][j+v] != 1 {
						add = false
						break
					}
					// 右边
					if matrix[i-v][j+expand] != 1 {
						add = false
						break
					}
				}
				if add {
					ans++
				} else {
					break
				}
			}
		}
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	matrix := Deserialize[[][]int](ReadLine(stdin))
	ans := countSquares(matrix)

	fmt.Println("\noutput:", Serialize(ans))
}
