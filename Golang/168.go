package main

import "fmt"

func convertToTitle(columnNumber int) string {
	result := ""
	if columnNumber == 0 {
		return result
	}
	list := []int{}
	if columnNumber <= 26 {
		result += string(columnNumber + 64)
		return result
	} else {
		for {
			if columnNumber >= 0 && columnNumber <= 26 {
				if columnNumber == 0 {
					list[len(list)-1] -= 1
					list = append(list, 26)
				} else {
					list = append(list, columnNumber)
				}
				break
			} else {
				if columnNumber%26 == 0 {
					list = append(list, 26)
					list = append(list, (columnNumber/26)-1)
					break
				} else {
					list = append(list, columnNumber%26)
				}
			}
			columnNumber = columnNumber / 26
		}
	}
	fmt.Println(list)
	for i := len(list) - 1; i >= 0; i-- {
		result += string(64 + list[i])
	}
	return result
}
