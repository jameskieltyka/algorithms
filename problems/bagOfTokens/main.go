package main

import "sort"

//https://leetcode.com/problems/bag-of-tokens

func bagOfTokensScore(tokens []int, P int) int {
	if len(tokens) == 0 {
		return 0
	}

	sort.Ints(tokens)

	point := 0
	i, j := 0, len(tokens)-1
	for i <= j {
		if P-tokens[i] >= 0 {
			P -= tokens[i]
			point++
			i++
		} else if point > 0 && i <= j-1 {
			P += tokens[j]
			j--
			point--
		} else {
			break
		}
	}

	return point
}
