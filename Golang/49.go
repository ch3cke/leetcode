package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	if strs == nil || len(strs) == 0 {
		return [][]string{}
	}
	tmpMap := map[string][]string{}
	tmpTable := []string{}
	for i := 0; i < len(strs); i++ {
		l := change(strs[i])
		tmpTable = append(tmpTable, l)
		if _, ok := tmpMap[l]; ok {
			tmpMap[l] = append(tmpMap[l], strs[i])
		} else {
			tmpMap[l] = []string{strs[i]}
		}
	}
	sort.Strings(tmpTable)
	result := [][]string{}
	result = append(result, tmpMap[tmpTable[0]])
	for j := 1; j < len(tmpTable); j++ {
		if tmpTable[j-1] != tmpTable[j] {
			result = append(result, tmpMap[tmpTable[j]])
		}
	}
	return result
}

func change(str string) string {
	tmp := []int{}
	sq := make([]byte, len(str))
	copy(sq, str)
	ss := string(sq)
	for i := 0; i < len(str); i++ {
		tmp = append(tmp, int(ss[i]))
	}
	sort.Ints(tmp)
	result := ""
	for j := 0; j < len(tmp); j++ {
		result += string(rune(tmp[j]))
	}
	return result
}
