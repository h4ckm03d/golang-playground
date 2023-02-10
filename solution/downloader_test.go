package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		A    []string
		want bool
	}{
		{
			name: "race",
			A:    []string{"http://www.example.com/foo/bar", "http://another.example.com/zot.zip", "http://www.example.com/foo/bar", "http://another.example.com/zot.zip"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
