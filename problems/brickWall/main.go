package main

import "fmt"

func main() {
	wall := [][]int{{1, 2, 2, 1},
		{3, 1, 2},
		{1, 3, 2},
		{2, 4},
		{3, 1, 2},
		{1, 3, 1, 1}}
	fmt.Println(leastBricks(wall))
}

func leastBricks(wall [][]int) int {

	joins := make(map[int]int)

	for i := range wall {
		pos := 0
		for j := 0; j < len(wall[i])-1; j++ {
			pos += wall[i][j]
			joins[pos]++
		}
	}

	maxJoin := 0
	for _, join := range joins {
		if join > maxJoin {
			maxJoin = join
		}
	}

	return len(wall) - maxJoin
}
