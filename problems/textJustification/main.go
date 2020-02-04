package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	for _, w := range fullJustify(words, 16) {
		fmt.Printf("\"%s\"\n", w)
	}
	fmt.Println()
	words = []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	for _, w := range fullJustify(words, 16) {
		fmt.Printf("\"%s\"\n", w)
	}
	fmt.Println()
	words = []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain",
		"to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	for _, w := range fullJustify(words, 20) {
		fmt.Printf("\"%s\"\n", w)
	}
}

func fullJustify(words []string, maxWidth int) []string {
	var res []string

	for len(words) > 0 {
		counter := 0
		i := 0
		for i < len(words) && counter+len(words[i]) <= maxWidth {
			counter += len(words[i]) + 1
			i++
		}

		var selected []string
		selected, words = words[0:i], words[i:len(words)]

		spaces := maxWidth + i - counter
		if len(words) > 0 {
			if i == 1 {
				res = append(res, selected[0]+strings.Repeat(" ", spaces))
			} else {
				equalSpaces := (spaces / (i - 1))
				remainingSpace := spaces % (i - 1)
				var s strings.Builder

				for j := 0; j < i-1; j++ {
					space := equalSpaces
					if remainingSpace > 0 {
						remainingSpace--
						space++
					}
					s.WriteString(selected[j])
					s.WriteString(strings.Repeat(" ", space))
				}
				s.WriteString(selected[i-1])
				res = append(res, s.String())
			}

		} else {
			res = append(res, strings.Join(selected, " ")+strings.Repeat(" ", spaces-i+1))
		}

	}
	return res
}
