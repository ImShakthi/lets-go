package basics_test

import (
	"lets-go/basics"
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "add all to get 6",
			args: args{
				numbers: []int{1, 2, 3},
			},
			want: 6,
		},
		{
			name: "add all to get 0",
			args: args{
				numbers: []int{-10, 2, -5, 12, 1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := basics.Add(tt.args.numbers...); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "get 120 for 5",
			args: args{
				n: 5,
			},
			want: 120,
		},
		{
			name: "get 1 for 0",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "get 1 for 1",
			args: args{
				n: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := basics.Factorial(tt.args.n); got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

