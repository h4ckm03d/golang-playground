package ann

import (
	"fmt"
	"testing"
)

func TestCreateNetwork(t *testing.T) {
	type args struct {
		name         string
		rate         float64
		input        Matrix
		output       Matrix
		hiddensNodes []int
	}
	tests := []struct {
		name  string
		args  args
		input []float64
		want  []float64
	}{
		{
			name: "Test case 1",
			args: args{
				name:         "XOR",
				rate:         0.1,
				input:        Matrix{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
				output:       Matrix{{0}, {1}, {1}, {0}},
				hiddensNodes: []int{50},
			},
			input: []float64{0, 1},
			want:  []float64{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := CreateNetwork(tt.args.name, tt.args.rate, tt.args.input, tt.args.output, tt.args.hiddensNodes...)
			n.Train(200)
			fmt.Println(n.Layers[len(n.Layers)-1])
			for i, v := range tt.args.input {
				fmt.Println(v, n.Predict(v), tt.args.output[i])
			}

			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("Transpose failed, excepted %v, got %v", tt.want, got)
			// }
		})
	}
}
