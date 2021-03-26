package main

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	var run func(i, gased int)
	result := false
	run = func(i, gased int) {
		a := gased + gas[i]
		b := i
		for {
			if a-cost[b] >= 0 {
				a = a - cost[b] + gas[(b+1)%n]
				b = (b + 1) % n
			} else {
				return
			}
			if b == i {
				result = true
				return
			}
		}
	}
	for i := 0; i < n; i++ {
		if !result {
			if gas[i]-cost[i] >= 0 {
				run((i+1)%n, gas[i]-cost[i])
			}
		} else {
			return i - 1
		}
	}
	if result {
		return n - 1
	} else {
		return -1
	}
}
