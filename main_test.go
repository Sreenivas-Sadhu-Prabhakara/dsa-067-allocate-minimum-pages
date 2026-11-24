package main

import "testing"

func TestAllocatePages(t *testing.T) {
	cases := []struct {
		pages    []int
		students int
		want     int
	}{
		{[]int{12, 34, 67, 90}, 2, 113},
		{[]int{10, 20, 30, 40}, 2, 60},
		{[]int{10, 20, 30, 40}, 4, 40},
		{[]int{10, 20, 30, 40}, 1, 100},
		{[]int{5}, 1, 5},
		{[]int{5}, 2, -1}, // more students than books
		{[]int{1, 2}, 0, -1},
	}
	for _, c := range cases {
		if got := AllocatePages(c.pages, c.students); got != c.want {
			t.Errorf("AllocatePages(%v,%d) = %d, want %d", c.pages, c.students, got, c.want)
		}
	}
}
