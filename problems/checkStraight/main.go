package main

import "fmt"

func main() {
	coordinates := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}
	fmt.Println(checkStraightLine(coordinates))

	coordinates = [][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}
	fmt.Println(checkStraightLine(coordinates))
}

func checkStraightLine(coordinates [][]int) bool {
	a2 := (coordinates[0][0] - coordinates[1][0])
	if a2 == 0 {
		for _, c := range coordinates {
			if c[1] != coordinates[0][1] {
				return false
			}
		}
		return true
	}

	a := (coordinates[0][1] - coordinates[1][1]) / a2
	b := coordinates[0][1] - a*coordinates[0][0]

	for _, c := range coordinates {
		if c[1] != a*c[0]+b {
			return false
		}
	}
	return true
}
