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
		input Matrix
	}{
		{
			name: "Test case: diabetic",
			args: args{
				name:         "diabetic",
				rate:         0.1,
				input:        Matrix{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}, {1, 1, 0}},
				output:       Matrix{{1}, {0}, {0}, {1}},
				hiddensNodes: []int{20},
			},
			input: Matrix{{0, 1, 1}, {1, 1, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := CreateNetwork(tt.args.name, tt.args.rate, tt.args.input, tt.args.output, tt.args.hiddensNodes...)
			n.Train(2000)
			fmt.Println(n.Layers[len(n.Layers)-1])
			for i, v := range tt.args.input {
				fmt.Println(v, n.Predict(v), tt.args.output[i])
			}

			for _, v := range tt.input {
				fmt.Println(v, n.Predict(v))
			}

			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("Transpose failed, excepted %v, got %v", tt.want, got)
			// }
		})
	}
}
