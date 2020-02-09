package main

func wordSets(dict []string, word string) [][]string {
	var res [][]string
	availableCombo(dict, word, []string{}, &res)
	if len(res) == 0 {
		return nil
	}
	return res

}

func availableCombo(dict []string, word string, currentWords []string, result *[][]string) {
	if word == "" {
		*result = append(*result, currentWords)
	}

	for _, option := range dict {

		if len(option) <= len(word) && word[0:len(option)] == option {
			currentWords = append(currentWords, option)
			availableCombo(dict, word[len(option):len(word)], currentWords, result)
			currentWords = currentWords[0 : len(currentWords)-1]
		}
	}
}
