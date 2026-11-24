package main

import "fmt"

// AllocatePages splits books (page counts, order fixed) among students so
// each student gets a contiguous run and the busiest student reads as few
// pages as possible. Returns that minimum, or -1 if students outnumber books.
func AllocatePages(pages []int, students int) int {
	if students <= 0 || students > len(pages) {
		return -1
	}
	lo, hi := 0, 0
	for _, p := range pages {
		if p > lo {
			lo = p // at least one student gets the thickest book
		}
		hi += p // one student reads everything
	}
	for lo < hi {
		mid := lo + (hi-lo)/2
		if studentsNeeded(pages, mid) <= students {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// studentsNeeded greedily packs books left to right under the cap.
func studentsNeeded(pages []int, cap int) int {
	count, cur := 1, 0
	for _, p := range pages {
		if cur+p > cap {
			count++
			cur = p
		} else {
			cur += p
		}
	}
	return count
}

func main() {
	fmt.Println(AllocatePages([]int{12, 34, 67, 90}, 2))
}
