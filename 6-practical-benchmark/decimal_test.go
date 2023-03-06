package main

import (
	"math/big"
	"testing"

	"github.com/shopspring/decimal"
	"go.mercari.io/go-bps/bps"
)

func decimalAdd(addition decimal.Decimal, n int) decimal.Decimal {
	result := decimal.NewFromInt(0)
	for i := 0; i < n; i++ {
		result = result.Add(addition)
	}

	return result
}

func bpsAdd(addition *bps.BPS, n int) *bps.BPS {
	result := bps.NewFromAmount(0)
	for i := 0; i < n; i++ {
		result = result.Add(addition)
	}
	return result
}

func BenchmarkShopspring001DecimalAdd(b *testing.B) {
	addition := decimal.NewFromFloat(0.01)
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		b.StartTimer()
		decimalAdd(addition, n)
	}
}

func BenchmarkShopspringPPMDecimalAdd(b *testing.B) {
	addition := decimal.NewFromFloat(0.000001)
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		b.StartTimer()
		decimalAdd(addition, n)
	}
}

func BenchmarkMercari001BpsAdd(b *testing.B) {
	addition := bps.NewFromPercentage(1)
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		b.StartTimer()
		bpsAdd(addition, n)
	}
}

func BenchmarkMercariPPMBpsAdd(b *testing.B) {
	addition := bps.NewFromPPM(big.NewInt(1))
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		b.StartTimer()
		bpsAdd(addition, n)
	}
}

func BenchmarkShopspringStringDecimalConvert(b *testing.B) {
	value := decimal.NewFromFloat(0.000001)
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		b.StartTimer()
		value.StringFixed(10)
	}
}

func BenchmarkMercariStringBpsConvert(b *testing.B) {
	value := bps.NewFromPPM(big.NewInt(1))
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		b.StartTimer()
		value.FloatString(10)
	}
}
