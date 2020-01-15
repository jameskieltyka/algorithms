package main

import (
	"fmt"
	"strings"
)

func main() {
	emails := []string{"test.email+alex@leetcode.com", "test.email.leet@code.com"}
	fmt.Println(numUniqueEmails(emails))
}

func numUniqueEmails(emails []string) int {
	unique := make(map[string]bool)
	for _, email := range emails {
		splitE := strings.Split(email, "@")
		local, domain := splitE[0], splitE[1]
		local = strings.Split(local, "+")[0]
		local = strings.Replace(local, ".", "", -1)
		unique[local+"@"+domain] = true
	}

	return len(unique)
}
