package main

//https://leetcode.com/problems/find-the-difference/submissions/

func findTheDifference(s string, t string) byte {
	letterMap := make(map[rune]int)

	for _, r := range s {
		letterMap[r]++
	}

	for _, r := range t {
		if i, ok := letterMap[r]; ok {
			if i == 0 {
				return byte(r)
			}
			letterMap[r]--
		} else {
			return byte(r)
		}
	}

	return 0
}

func findTheDifferenceXOR(s string, t string) byte {
	var c byte
	for i, _ := range s {
		c ^= s[i]
	}

	for i, _ := range t {
		c ^= t[i]
	}

	return c
}
