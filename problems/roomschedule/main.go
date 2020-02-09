package main

import (
	"sort"
)

//Given an array of time intervals (start, end) for classroom lectures (possibly overlapping), find the minimum number of rooms required.
//For example, given [(30, 75), (0, 50), (60, 150)], you should return 2.

type booking struct {
	start int
	end   int
}

func minimumRooms(bookings []booking) int {
	start := make([]int, len(bookings))
	end := make([]int, len(bookings))
	for i, b := range bookings {
		start[i] = b.start
		end[i] = b.end
	}

	sort.Ints(start)
	sort.Ints(end)

	endPos := 0
	count, maxCount := 0, 0
	for i := range start {
		if count > maxCount {
			maxCount = count
		}
		if start[i] < end[endPos] {
			count++
		} else {
			for end[endPos] < start[i] {
				endPos++
				count--
			}
			count++
		}
	}
	return maxCount
}
