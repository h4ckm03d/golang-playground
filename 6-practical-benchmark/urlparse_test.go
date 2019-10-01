package main

import (
	"net/url"
	"testing"
)

func BenchmarkParse(b *testing.B) {
	testUrl := "https://www.example.com/path/file.html?param1=value1&param2=123"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := url.Parse(testUrl)
		if err != nil {
			panic(err)
		}
	}
}
