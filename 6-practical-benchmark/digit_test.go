package main

import (
	fmt "fmt"
	"testing"
)

func sumUseInt(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit
		n /= 10
	}
	return sum
}

func sumUseStr(n int) int {
	sum := 0
	for _, r := range fmt.Sprintf("%d", n) {
		sum += int(r - '0')
	}
	return sum
}

func BenchmarkSumUseInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sumUseInt(n)
	}
}

func BenchmarkSumUseStr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sumUseStr(n)
	}
}
