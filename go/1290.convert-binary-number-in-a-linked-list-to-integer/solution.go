// Created by cubatic45 at 2025/07/14 17:51
// leetgo: 1.4.14
// https://leetcode.cn/problems/convert-binary-number-in-a-linked-list-to-integer/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func getDecimalValue(head *ListNode) (ans int) {
	for head != nil {
		ans = ans<<1 | head.Val
		// ans = ans<<1 + head.Val
		// ans = ans*2 + head.Val
		head = head.Next
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	ans := getDecimalValue(head)

	fmt.Println("\noutput:", Serialize(ans))
}
