package main

import "strconv"

func fractionToDecimal(numerator int, denominator int) string {
	result := map[int]int{}
	numerator2 := 0
	flag := false
	s := ""
	index := []int{}
	if numerator == 0 {
		return "0"
	}

	if numerator*denominator < 0 {
		s += "-"
		if numerator < 0 {
			numerator = 0 - numerator
		}
		if denominator < 0 {
			denominator = 0 - denominator
		}
	} else {
		if numerator < 0 {
			numerator = 0 - numerator
		}
		if denominator < 0 {
			denominator = 0 - denominator
		}
	}
	if denominator <= numerator {
		tmp := numerator / denominator
		numerator = numerator % denominator
		if numerator == 0 {
			return s + strconv.Itoa(tmp)
		}
		s += (strconv.Itoa(tmp) + ".")
	} else {
		s += "0."
	}
	for {
		numerator = numerator * 10
		for numerator < denominator {
			_, ok := result[numerator]
			if ok {
				flag = true
				break
			} else {
				index = append(index, numerator)
				result[numerator] = 0
			}
			numerator = numerator * 10
		}
		if flag {
			break
		}
		tmp := numerator / denominator
		numerator2 = numerator % denominator
		if numerator2 == 0 {
			for i := 0; i < len(index); i++ {
				s += strconv.Itoa(result[index[i]])
			}
			return s + strconv.Itoa(tmp)
		}
		_, ok := result[numerator]
		if ok {
			break
		} else {
			index = append(index, numerator)
			result[numerator] = tmp
		}
		numerator = numerator2
	}
	for i := 0; i < len(index); i++ {
		if index[i] == numerator {
			s += "("
		}
		s += strconv.Itoa(result[index[i]])
	}
	return s + ")"
}
