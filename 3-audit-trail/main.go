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
		r.diffs = append(r.diffs, fmt.Sprintf("{\"field\":\"%s\", \"old\": \"%+v\", \"new\": \"%+v\"}", r.path.String(), vx, vy))
	}
}

func (r *DiffReporter) PopStep() {
	r.path = r.path[:len(r.path)-1]
}

func (r *DiffReporter) String() string {
	return "{" + strings.Join(r.diffs, ",\n") + "}"
}

type Person struct {
	ID        uint      `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func main() {
	p1 := Person{1, "Lutfi", time.Now()}
	p2 := Person{1, "Moch", time.Now()}
	var r DiffReporter

	if !cmp.Equal(p1, p2, cmp.Reporter(&r)) {
		fmt.Print(r.String())
	}
}
