package main

import "fmt"

//Come back to this valid-permutations-for-di-sequence/

func main() {
	fmt.Println(numPermsDISequence("IDDDII"))
}

var mod int = 1e9 + 7

func numPermsDISequence(S string) int {
	n := len(S)
	cache := make(map[string]int)
	available := make([]bool, n+1)
	for i := range available {
		available[i] = true
	}

	count := 0
	for i := range available {
		available[i] = false
		count += permuation(S, 0, i, available, cache)
		available[i] = true
	}

	return count
}

func permuation(S string, pos int, prev int, options []bool, cache map[string]int) int {
	if pos == len(S) {
		return 1
	}

	key := fmt.Sprintf("%d%d%v", pos, prev, options)
	if n, ok := cache[key]; ok {
		return n
	}

	count := 0
	if S[pos] == 'I' {
		for i := prev + 1; i < len(options); i++ {
			if options[i] {
				options[i] = false
				count += (permuation(S, pos+1, i, options, cache) % mod)
				options[i] = true
			}
		}
	} else {
		for i := prev - 1; i >= 0; i-- {
			if options[i] {
				options[i] = false
				count += (permuation(S, pos+1, i, options, cache) % mod)
				options[i] = true
			}
		}
	}

	cache[key] = count
	return count
}
