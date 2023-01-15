package main

import (
	fmt "fmt"
	"testing"
)

func Solution1(n int) int {
	sum := 0
	for _, r := range fmt.Sprintf("%d", n) {
		sum += int(r - '0')
	}
	return sum
}

func Solution2(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit
		n /= 10
	}
	return sum
}

func BenchmarkSumUseInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Solution2(n)
	}
}

func BenchmarkSumUseStr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Solution1(n)
	}
}
