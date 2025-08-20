// Created by cubatic45 at 2025/08/19 17:55
// leetgo: 1.4.14
// https://leetcode.cn/problems/number-of-zero-filled-subarrays/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

/*
0 --> 1
0,0 --> 1+2
0,0,0 --> 3+3 f2 + len
0,0,0,0 --> 6+4 f3 + len
f(n+1)=f(n)+n+1
*/

var m = []int64{0, 1, 3}

func s(x int64) int64 {
	if x < int64(len(m)) {
		return m[x]
	}
	for i := int64(len(m)); i <= x; i++ {
		m = append(m, m[i-1]+int64(i))
	}
	return m[x]
}

func zeroFilledSubarray(nums []int) (ans int64) {
	var cnt int64
	for _, v := range nums {
		if v == 0 {
			cnt++
			ans += cnt // 等价于 ans += s(cnt) - s(cnt-1)
		} else {
			cnt = 0
		}
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := zeroFilledSubarray(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
