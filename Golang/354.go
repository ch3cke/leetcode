package main

import (
	"fmt"
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		flag := false
		//return envelopes[i][0]<envelopes[j][0]
		if envelopes[i][0] < envelopes[j][0] {
			flag = true
		} else if envelopes[i][0] == envelopes[j][0] {
			if envelopes[i][1] > envelopes[j][1] {
				flag = true
			}
		}
		return flag
	})
	fmt.Println(envelopes)
	f := []int{}
	for _, e := range envelopes {
		h := e[1]
		if i := sort.SearchInts(f, h); i < len(f) {
			f[i] = h
		} else {
			f = append(f, h)
		}
		fmt.Println(f)
	}
	return len(f)
}
