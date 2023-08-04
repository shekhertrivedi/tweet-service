package service

import "testing"

func TestExampleFunc(t *testing.T) {
	type args struct {
		a  int
		b  int
		op string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "add",
			args: args{
				a:  10,
				b:  5,
				op: "add",
			},
			want: 15,
		},
		{
			name: "sub",
			args: args{
				a:  10,
				b:  5,
				op: "sub",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExampleFunc(tt.args.a, tt.args.b, tt.args.op); got != tt.want {
				t.Errorf("ExampleFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
