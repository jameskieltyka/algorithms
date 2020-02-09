package main

import (
	"strings"
)

//Suppose we represent our file system by a string in the following manner:
//"dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"
//Given a string representing the file system in the above format, return the length of the longest absolute path to a file in the abstracted file system. If there is no file in the system, return 0.

func longestPath(fs string) int {

	res := 0
	maxPath := 0
	prevLevel := 0
	files := strings.Split(fs, "\n")
	levelAdded := make([]int, len(files))
	for i := range files {
		index := 0
		level := 0
		for files[i][index:index+1] == "\t" {
			level++
			index = index + 1
		}

		if level <= prevLevel {
			for j := prevLevel; j >= level; j-- {
				maxPath -= levelAdded[j]
				levelAdded[j] = 0
			}
		}

		levelAdded[level] = len(files[i]) - index + 1
		maxPath += levelAdded[level]
		if strings.Contains(files[i], ".") {
			if res < maxPath {
				res = maxPath
			}
		}
		prevLevel = level
	}
	if res > 0 {
		return res - 1
	}
	return 0
}
