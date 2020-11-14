package main

func divide(dividend int, divisor int) int {
	result := 0
	add := 1
	addnum := 0
	check2 := 0
	if divisor < 0 {
		check2 = 0 - divisor
	} else {
		check2 = divisor
	}
	recode := check2
	check := 0
	if dividend > 0 {
		check = dividend
	} else {
		check = 0 - dividend
	}

	if check2 == 0 {
		return 0
	}
	if check < check2 {
		return 0
	}
	if check == check2 {
		if (divisor > 0 && dividend > 0) || (divisor < 0 && dividend < 0) {
			return 1
		} else {
			return -1
		}
	}

	if check2 == 1 {
		if (divisor > 0 && dividend > 0) || (divisor < 0 && dividend < 0) {
			if check > 2147483647 {
				return 2147483647
			} else {
				return check
			}
		} else {
			return -check
		}
	}

	for {
		addnum = addnum + recode
		result = result + add
		if addnum > check {
			result = result - add
			add = 1
			addnum = addnum - recode
			recode = check2
			if addnum+check2 > check {
				if (divisor > 0 && dividend > 0) || (divisor < 0 && dividend < 0) {
					if result > 2147483647 {
						return 2147483647
					} else {
						return result
					}
				} else {
					return 0 - result
				}
			}
			continue
		}
		recode = recode + recode
		add = add + add
	}
}
