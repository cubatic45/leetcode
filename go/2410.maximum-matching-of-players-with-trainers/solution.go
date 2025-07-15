// Created by cubatic45 at 2025/07/14 18:02
// leetgo: 1.4.14
// https://leetcode.cn/problems/maximum-matching-of-players-with-trainers/

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func matchPlayersAndTrainers(players []int, trainers []int) (ans int) {
	if len(players) == 0 || len(trainers) == 0 {
		return
	}
	slices.Sort(players)
	slices.Sort(trainers)
	i, j := 0, 0
	for {
		if i >= len(players) || j >= len(trainers) {
			return
		}
		if players[i] <= trainers[j] {
			ans++
			j++
			i++
		} else {
			j++
		}
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	players := Deserialize[[]int](ReadLine(stdin))
	trainers := Deserialize[[]int](ReadLine(stdin))
	ans := matchPlayersAndTrainers(players, trainers)

	fmt.Println("\noutput:", Serialize(ans))
}
