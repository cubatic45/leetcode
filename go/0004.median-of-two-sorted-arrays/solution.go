// Created by cubatic45 at 2025/02/14 11:29
// leetgo: 1.4.13
// https://leetcode.cn/problems/median-of-two-sorted-arrays/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findMedianSortedArrays(nums1 []int, nums2 []int) (ans float64) {
	l1Len := len(nums1)
	l2Len := len(nums2)

	// 双指针合并两个有序数组，找到中位数
	i, j := 0, 0
	index := (l1Len + l2Len) / 2
	isOdd := (l1Len+l2Len)%2 == 1
	var prev, curr int

	for k := 0; k <= index; k++ {
		prev = curr
		if i < l1Len && (j >= l2Len || nums1[i] <= nums2[j]) {
			curr = nums1[i]
			i++
		} else {
			curr = nums2[j]
			j++
		}
	}
    

	if isOdd {
		return float64(curr)
	}

	// 如果总数为偶数，返回中位数的平均值
	return float64(prev+curr) / 2.0
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums1 := Deserialize[[]int](ReadLine(stdin))
	nums2 := Deserialize[[]int](ReadLine(stdin))
	ans := findMedianSortedArrays(nums1, nums2)

	fmt.Println("\noutput:", Serialize(ans))
}
