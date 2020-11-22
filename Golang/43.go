package main

import (
	"fmt"
	"math"
	"math/big"
)

func multiply(num1 string, num2 string) string {
	ip1 := big.NewInt(0)
	tmp := big.NewInt(10)
	for i := 0; i < len(num1); i++ {
		ip1.Mul(ip1, tmp).Add(ip1, big.NewInt(int64(num1[i]-48)))
		fmt.Println(ip1)
	}
	ip2 := big.NewInt(0)
	tmp2 := big.NewInt(10)
	for i := 0; i < len(num2); i++ {
		ip2.Mul(ip2, tmp2).Add(ip2, big.NewInt(int64(num2[i]-48)))
		fmt.Println(ip2)
	}
	result := big.NewInt(math.MaxInt64)
	result.Mul(ip1, ip2)

	return result.String()
}
