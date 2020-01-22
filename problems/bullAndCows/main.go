package main

import "fmt"

func main() {
	fmt.Println(getHint("1807", "7810"))
	fmt.Println(getHint("1123", "0111"))
	fmt.Println(getHint("1201", "1100"))
	fmt.Println(getHint("1122", "1222"))
	fmt.Println(getHint("6244988279", "3819888600"))
}

func getHint(secret string, guess string) string {
	charMap := make(map[rune]int)
	for _, c := range secret {
		charMap[c]++
	}

	correct, missed := 0, 0
	for i, c := range guess {
		if secret[i] == guess[i] {
			correct++
			charMap[c]--
			if charMap[c] < 0 {
				missed += charMap[c]
				charMap[c] = 0
			}
		} else if _, ok := charMap[c]; ok {
			if charMap[c] > 0 {
				charMap[c]--
				missed++
			}
		}
	}

	return fmt.Sprintf("%dA%dB", correct, missed)
}
