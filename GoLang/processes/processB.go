package processes

import (
	"fmt"
)

func ProcessB(fileData []byte) {
	// Reverse the string
	reversedString := reverseString(string(fileData))
	fmt.Printf("Processing file type B: Original string: %s, Reversed string: %s\n", string(fileData), reversedString)
	fmt.Println()
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
