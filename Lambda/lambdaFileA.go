package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/Knetic/govaluate"
)

func main() {
	// Set your AWS region
	region := "us-east-1"

	// Initialize AWS SDK configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		fmt.Printf("Error loading AWS SDK config: %v\n", err)
		return
	}

	// Create an S3 client
	client := s3.NewFromConfig(cfg)

	// Specify your S3 bucket and folder
	bucket := "awscloudfinalproject2024"
	folder := "Type A/"

	// List objects in the specified folder
	listInput := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(folder),
	}

	listOutput, err := client.ListObjectsV2(context.TODO(), listInput)
	if err != nil {
		fmt.Printf("Error listing objects in S3: %v\n", err)
		return
	}

	// Process each object in the folder
	for _, object := range listOutput.Contents {
		// Download the object's content
		getInput := &s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    object.Key,
		}

		getOutput, err := client.GetObject(context.TODO(), getInput)
		if err != nil {
			fmt.Printf("Error getting object from S3: %v\n", err)
			return
		}

		defer getOutput.Body.Close()

		// Read the content of the object
		fileData, err := ioutil.ReadAll(getOutput.Body)
		if err != nil {
			fmt.Printf("Error reading file data: %v\n", err)
			return
		}

		// Call the processing algorithm
		result := ProcessA(fileData)

		// Write the result back to the same file
		err = ioutil.WriteFile(fmt.Sprintf("/tmp/%s", *object.Key), []byte(result), 0644)
		if err != nil {
			fmt.Printf("Error writing result to file: %v\n", err)
			return
		}
	}
}

func ProcessA(fileData []byte) string {
	fmt.Print("Processing file type A:", string(fileData))

	expression, err := govaluate.NewEvaluableExpression(string(fileData))
	if err != nil {
		fmt.Printf("Error creating expression: %v\n", err)
		return ""
	}
	result, err := expression.Evaluate(nil)
	if err != nil {
		fmt.Printf("Error evaluating expression: %v\n", err)
		return ""
	}

	resultString := fmt.Sprintf("Result of expression: %v\n\n", result)
	fmt.Print(resultString)

	return resultString
}