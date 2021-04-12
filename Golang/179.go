package main

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	ss := make([]string, len(nums))
	for i, num := range nums {
		ss[i] = strconv.Itoa(num)
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i]+ss[j] >= ss[j]+ss[i]
	})
	result := strings.Join(ss, "")
	if result[0] == '0' {
		return "0"
	}
	return result
}
