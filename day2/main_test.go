package main

import "testing"

func Test_runIntCodeProgram(t *testing.T) {
	type args struct {
		ints []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}},
			want: 3500,
		}, {
			name: "Example 2",
			args: args{[]int{1, 0, 0, 0, 99}},
			want: 2,
		}, {
			name: "Example 2",
			args: args{[]int{2, 3, 0, 3, 99}},
			want: 2,
		}, {
			name: "Example 3",
			args: args{[]int{2, 4, 4, 5, 99, 0}},
			want: 2,
		}, {
			name: "Example 4",
			args: args{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runIntCodeProgram(tt.args.ints); got != tt.want {
				t.Errorf("runIntCodeProgram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findNounAndVerb(t *testing.T) {
	type args struct {
		ints   []int
		output int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				ints:   readFile("./input.txt"),
				output: 3085697,
			},
			want: 1202,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNounAndVerb(tt.args.ints, tt.args.output); got != tt.want {
				t.Errorf("findNounAndVerb() = %v, want %v", got, tt.want)
			}
		})
	}
}
