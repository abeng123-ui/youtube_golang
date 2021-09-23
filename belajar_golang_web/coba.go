package main

import (
	"fmt"
)

func main() {
	displayMulti("Jakarta", 2)
}

func displayMulti(kata string, n int) {
	total := len(kata)
	var runes []rune

	if total < 3 {
		runes = []rune(reverseString(kata))
	} else {
		runes = []rune(kata)
	}
	safeSubstring := string(runes[0:3])

	for i := 0; i < n; i++ {
		fmt.Print(safeSubstring)
	}

}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
