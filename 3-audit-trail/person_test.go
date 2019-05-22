package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/bxcodec/faker"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
	"github.com/haritsfahreza/libra"
	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	t.Run("Compare person", func(t *testing.T) {
		p1 := Person{1, "Lutfi", Job{ID: 1}, time.Now()}
		p2 := Person{1, "Moch", Job{ID: 2, Name: "Programmer"}, time.Now()}
		diff, err := Compare(p1, p2)
		assert.Nil(t, err)
		spew.Dump(diff)
	})

	t.Run("Compare job", func(t *testing.T) {
		p1 := Job{1, "oke", 1000, time.Now()}
		p2 := Job{2, "oke", 10000, time.Now()}
		diff, err := Compare(p1, p2)
		assert.Nil(t, err)
		spew.Dump(diff)
	})

	t.Run("Compare job with person", func(t *testing.T) {
		p1 := Job{1, "oke", 1000, time.Now()}
		p2 := Person{2, "oke", Job{}, time.Now()}
		diff, err := Compare(p1, p2)
		assert.NotNil(t, err)
		spew.Dump(diff)
	})
}

func BenchmarkCompare(b *testing.B) {
	p1 := Person{1, "Lutfi", Job{ID: 1}, time.Now()}
	p2 := Person{1, "Moch", Job{ID: 1}, time.Now()}
	for n := 0; n < b.N; n++ {
		Compare(p1, p2)
	}
}

func BenchmarkCmpDiff(b *testing.B) {
	p1 := Person{1, "Lutfi", Job{ID: 1}, time.Now()}
	p2 := Person{1, "Moch", Job{ID: 1}, time.Now()}
	for n := 0; n < b.N; n++ {
		cmp.Diff(p1, p2)
	}
}

func BenchmarkLibra(b *testing.B) {
	p1 := Person{1, "Lutfi", Job{ID: 1}, time.Now()}
	p2 := Person{1, "Moch", Job{ID: 1}, time.Now()}
	ctx := context.Background()
	for n := 0; n < b.N; n++ {
		libra.Compare(ctx, p1, p2)
	}
}

func createJobs() []Job {
	N := 1000000
	jobs := make([]Job, N)
	for i := range jobs {
		err := faker.FakeData(&jobs[i])
		if err != nil {
			fmt.Println(err)
		}
	}
	return jobs
}
