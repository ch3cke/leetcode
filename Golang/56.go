package main

func merge(intervals [][]int) [][]int {
	sort(intervals)
	for i := 0; i < len(intervals)-1; {
		if intervals[i][1] >= intervals[i+1][0] {
			if intervals[i][1] < intervals[i+1][1] {
				intervals[i][1] = intervals[i+1][1]
			}
			intervals = append(intervals[:i+1], intervals[i+2:]...)
		} else {
			i++
		}
	}
	return intervals
}

func sort(intervals [][]int) {
	for i := 0; i < len(intervals); i++ {
		for j := i; j < len(intervals); j++ {
			if intervals[i][0] > intervals[j][0] {
				tmp := intervals[j]
				intervals[j] = intervals[i]
				intervals[i] = tmp
			}
		}
	}
}
