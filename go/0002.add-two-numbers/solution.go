// Created by cubatic45 at 2025/02/13 18:24
// leetgo: 1.4.13
// https://leetcode.cn/problems/add-two-numbers/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (ans *ListNode) {
	i, j, add, ans := 0, 0, 0, new(ListNode)
	l := ans
	for {
		if l1 == nil && l2 == nil && add == 0 {
			break
		}
		if l1 == nil {
			i = 0
		} else {
			i = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			j = 0
		} else {
			j = l2.Val
			l2 = l2.Next
		}
		mod := (i + j + add) % 10
		add = (i + j + add) / 10
		l.Val = mod
		if l1 != nil || l2 != nil || add != 0 {
			l.Next = new(ListNode)
			l = l.Next
		}
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	l1 := Deserialize[*ListNode](ReadLine(stdin))
	l2 := Deserialize[*ListNode](ReadLine(stdin))
	ans := addTwoNumbers(l1, l2)

	fmt.Println("\noutput:", Serialize(ans))
}
