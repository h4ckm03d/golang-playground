package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/go-cmp/cmp"
)

// DiffReporter is a simple custom reporter that only records differences
// detected during comparison.
type DiffReporter struct {
	path  cmp.Path
	diffs []string
}

func (r *DiffReporter) PushStep(ps cmp.PathStep) {
	r.path = append(r.path, ps)
}

func (r *DiffReporter) Report(rs cmp.Result) {
	if !rs.Equal() {
		vx, vy := r.path.Last().Values()
		r.diffs = append(r.diffs, fmt.Sprintf("{\"field\":\"%v\", \"old\": \"%+v\", \"new\": \"%+v\"}", r.path, vx, vy))
	}
}

func (r *DiffReporter) PopStep() {
	r.path = r.path[:len(r.path)-1]
}

func (r *DiffReporter) String() string {
	return "{" + strings.Join(r.diffs, ",\n") + "}"
}

func main() {
	p1 := Person{1, "Lutfi", time.Now()}
	p2 := Person{1, "Moch", time.Now()}
	var r DiffReporter

	if !cmp.Equal(p1, p2, cmp.Reporter(&r)) {
		fmt.Print(r.String())
	}
	a := map[string]interface{}{
		"b": 1,
		"c": "asc",
	}

	b := map[string]interface{}{
		"b": 12,
		"c": "asc",
	}

	if diff := cmp.Diff(a, b); diff != "" {
		fmt.Print(diff)
	}
}
