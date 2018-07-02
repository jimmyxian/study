package main

import (
	"fmt"
)

func getDP(array []int) []int {
	rst := make([]int, len(array))
	for i := 0; i < len(array); i++ {
		rst[i] = 0
	}

	for j := 0; j < len(array); j++ {
		for i := 0; i < j; i++ {
			if array[i] < array[j] && rst[j] < rst[i]+1 {
				rst[j] = rst[i] + 1
			}
		}
	}
	return rst
}

func findMax(array []int) {
	dp := getDP(array)

	max := 0
	for _, v := range dp {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)
}

func GetMax(array []int) {
	dp := getDP(array)

	maxindex := 0
	maxlen := 0
	for k, v := range dp {
		if v > maxlen {
			maxlen = v
			maxindex = k
		}
	}

	rst := make([]int, maxlen)
	rst[maxlen-1] = array[maxindex]
	maxlen--
	for i := maxindex - 1; i > 0; i-- {
		if array[i] < array[maxindex] && dp[i] == dp[maxindex]-1 {
			maxlen--
			rst[maxlen] = array[i]
			maxindex = i
		}
	}

	fmt.Println(rst)
}

func main() {
	array := []int{1, 4, 5, 6, 2, 3, 8}
	GetMax(array)
}
