// package Lambda

// import (
// 	"context"
// 	"fmt"
// 	"io/ioutil"
// 	"main/processes"
// 	"github.com/aws/aws-sdk-go-v2/aws"
// 	"github.com/aws/aws-sdk-go-v2/config"
// 	"github.com/aws/aws-sdk-go-v2/service/s3"
// )

// func ProcessC() {
// 	// Set your AWS region
// 	region := "US East (N. Virginia) us-east-1"

// 	// Initialize AWS SDK configuration
// 	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
// 	if err != nil {
// 		fmt.Printf("Error loading AWS SDK config: %v\n", err)
// 		return
// 	}

// 	// Create an S3 client
// 	client := s3.NewFromConfig(cfg)

// 	// Specify your S3 bucket and folder
// 	bucket := "awscloudfinalproject2024"
// 	folder := "Type C/"

// 	// List objects in the specified folder
// 	listInput := &s3.ListObjectsV2Input{
// 		Bucket: aws.String(bucket),
// 		Prefix: aws.String(folder),
// 	}

// 	listOutput, err := client.ListObjectsV2(context.TODO(), listInput)
// 	if err != nil {
// 		fmt.Printf("Error listing objects in S3: %v\n", err)
// 		return
// 	}

// 	// Process each object in the folder
// 	for _, object := range listOutput.Contents {
// 		// Download the object's content
// 		getInput := &s3.GetObjectInput{
// 			Bucket: aws.String(bucket),
// 			Key:    object.Key,
// 		}

// 		getOutput, err := client.GetObject(context.TODO(), getInput)
// 		if err != nil {
// 			fmt.Printf("Error getting object from S3: %v\n", err)
// 			return
// 		}

// 		defer getOutput.Body.Close()

// 		// Read the content of the object
// 		fileData, err := ioutil.ReadAll(getOutput.Body)
// 		if err != nil {
// 			fmt.Printf("Error reading file data: %v\n", err)
// 			return
// 		}

// 		// Call the processing algorithm
// 		processes.ProcessC(fileData)
// 	}
// }