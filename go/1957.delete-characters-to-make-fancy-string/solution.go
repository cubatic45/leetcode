// Created by cubatic45 at 2025/07/22 17:09
// leetgo: 1.4.14
// https://leetcode.cn/problems/delete-characters-to-make-fancy-string/

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"unsafe"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func makeFancyString(s string) string {
	if len(s) < 3 {
		return s
	}
	buf := new(bytes.Buffer)
	var lastByte byte
	var count int
	for _, v := range string2Bytes(s) {
		if v == lastByte {
			count++
		} else {
			lastByte = v
			count = 1
		}
		if count <= 2 {
			buf.WriteByte(v)
		}
	}
	return buf.String()
}

func string2Bytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := makeFancyString(s)

	fmt.Println("\noutput:", Serialize(ans))
}
