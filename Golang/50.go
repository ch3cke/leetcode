package main

func myPow(x float64, n int) float64 {
	if n > 0 {
		return mypowa(x, n)
	} else {
		return 1 / mypowa(x, -n)
	}
}

func mypowa(x float64, n int) float64 {
	N := n
	ans := 1.0
	tmpN := x
	for N > 0 {
		if N%2 == 1 {
			ans *= tmpN
		}
		tmpN *= tmpN
		N = N / 2
	}
	return ans
}
