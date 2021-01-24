package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	result := make(map[string]int)
	for _, w := range(words){
		result[w] += 1
	}
	return result
}

func main() {
	wc.Test(WordCount)
}

