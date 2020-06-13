package main

import "testing"

func TestCAS(t *testing.T) {

	var testCases = []struct {
		input  []int
		target int
		want   int
	}{
		{[]int{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}, 5, 1},
		{[]int{30, 40, 50, 10, 20}, 60, -1},
		{[]int{}, 1, -1},
		{[]int{4, 5, 10, 7, 8, 9, 0, 1, 2, 3}, 10, -1},
		{[]int{1, 2, 3, 4, 5, 6, 10, 11}, 10, 6},
	}

	for _, test := range testCases {
		got := CircularArraySearch(test.input, test.target)
		if got != test.want {
			t.Errorf("%v, got: %d, want: %d", test.input, got, test.want)
		}
	}
}
