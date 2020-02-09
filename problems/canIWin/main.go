package main

import "fmt"

//https://leetcode.com/problems/can-i-win

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if desiredTotal < maxChoosableInteger {
		return true
	}

	if (1+maxChoosableInteger)*maxChoosableInteger/2 < desiredTotal {
		return false
	}

	memo := make(map[string]bool)
	used := make([]bool, maxChoosableInteger)
	return turn(used, desiredTotal, memo)
}

func turn(used []bool, desired int, memo map[string]bool) bool {
	key := fmt.Sprintf("%v", used)
	if r, ok := memo[key]; ok {
		return r
	}

	for i := len(used) - 1; i+1 >= desired; i-- {
		if !used[i] {
			return true
		}
	}

	for i := range used {
		if !used[i] {
			used[i] = true
			if !turn(used, desired-(i+1), memo) {
				used[i] = false
				return true
			}
			used[i] = false
		}
	}

	memo[key] = false
	return false
}
