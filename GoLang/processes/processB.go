package processes

import (
	"fmt"
	"github.com/Knetic/govaluate"
)

func ProcessB(fileData []byte) {
	fmt.Print("Processing file type A:", string(fileData))

	expression, err := govaluate.NewEvaluableExpression(string(fileData))
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