package main

func generateMatrix(n int) [][]int {
	//init
	result := [][]int{}
	for i := 0; i < n; i++ {
		tmp := []int{}
		for j := 0; j < n; j++ {
			tmp = append(tmp, 0)
		}
		result = append(result, tmp)
	}
	a := 0
	b := n - 1
	c := 0
	d := n - 1
	tmp_num := n * n
	nums := 1
	for nums <= tmp_num {
		for i := a; i <= b; i++ {
			result[c][i] = nums
			nums++
		}
		c++
		for i := c; i <= b; i++ {
			result[i][b] = nums
			nums++
		}
		b--
		for i := b; i >= a; i-- {
			result[d][i] = nums
			nums++
		}
		d--
		for i := d; i >= c; i-- {
			result[i][a] = nums
			nums++
		}
		a++
	}
	return result
}
