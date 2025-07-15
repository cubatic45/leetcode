// Created by cubatic45 at 2025/07/15 09:21
// leetgo: 1.4.14
// https://leetcode.cn/problems/valid-word/

package main

import (
	"bufio"
	"fmt"
	"os"
	"unsafe"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func string2Bytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func isValid(word string) bool {
	yl := 0
	nl := 0
	l := len(word)

	for _, w := range string2Bytes(word) {
		switch {
		case w >= 'A' && w <= 'z':
			switch w {
			case 'a', 'e', 'i', 'o', 'u',
				'A', 'E', 'I', 'O', 'U':
				yl++
			}
		case w >= '0' && w <= '9':
			nl++
		default:
			return false
		}
	}
	return l >= 3 && yl >= 1 && l-yl-nl >= 1
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	word := Deserialize[string](ReadLine(stdin))
	ans := isValid(word)

	fmt.Println("\noutput:", Serialize(ans))
}
