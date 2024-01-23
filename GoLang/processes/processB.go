package processes

import (
	"fmt"
	"github.com/Knetic/govaluate"
)

func ProcessB(fileData []byte) {
	// Reverse the string
	reversedString := reverseString(string(fileData))
	fmt.Printf("Processing file type B. Original string: %s, Reversed string: %s\n", string(fileData), reversedString)

	// Pass the reversed string to govaluate
	expression, err := govaluate.NewEvaluableExpression(reversedString)
	if err != nil {
		fmt.Printf("Error creating expression: %v\n", err)
		return
	}

	result, err := expression.Evaluate(nil)
	if err != nil {
		fmt.Printf("Error evaluating expression: %v\n", err)
		return
	}

	fmt.Printf("Result of expression: %v\n\n", result)
}

// reverseString reverses a string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
