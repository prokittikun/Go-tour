package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	sSplit := strings.Fields(s)
	for i := 0; i < len(sSplit); i++ {
		m[sSplit[i]] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
