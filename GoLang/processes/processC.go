package processes

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"strings"
)

func ProcessC(fileData []byte) {
	fmt.Print("Processing file type C:", string(fileData))

	// Convert ASCII to binary
	binaryString := asciiToBinary(string(fileData))
	fmt.Printf("Binary representation: %s\n", binaryString)

	// Pass the binary string to govaluate
	expression, err := govaluate.NewEvaluableExpression(binaryString)
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

// asciiToBinary converts an ASCII string to binary representation
func asciiToBinary(s string) string {
	var binaryStrings []string
	for _, char := range s {
		binaryStrings = append(binaryStrings, fmt.Sprintf("%08b", char))
	}
	return strings.Join(binaryStrings, "")
}
