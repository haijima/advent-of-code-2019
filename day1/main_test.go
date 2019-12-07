package main

import "testing"

func Test_calculateTotalFuel(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{filepath: "test.txt"},
			want: 34241,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateTotalFuel(tt.args.filepath); got != tt.want {
				t.Errorf("calculateTotalFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateFuel1(t *testing.T) {
	type args struct {
		mass int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{mass: 12},
			want: 2,
		}, {
			name: "Example 2",
			args: args{mass: 14},
			want: 2,
		}, {
			name: "Example 3",
			args: args{mass: 1969},
			want: 654,
		}, {
			name: "Example 4",
			args: args{mass: 100756},
			want: 33583,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateFuel1(tt.args.mass); got != tt.want {
				t.Errorf("calculateFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateFuel2(t *testing.T) {
	type args struct {
		mass int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{14},
			want: 2,
		}, {
			name: "Example 2",
			args: args{1969},
			want: 966,
		}, {
			name: "Example 3",
			args: args{100756},
			want: 50346,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateFuel2(tt.args.mass); got != tt.want {
				t.Errorf("calculateFuel2() = %v, want %v", got, tt.want)
			}
		})
	}
}
