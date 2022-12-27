package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/smithy-go"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var functionName string
	flag.StringVar(&functionName, "function-name", "", "Name of the AWS Lambda function to update")
	var zipFile string
	flag.StringVar(&zipFile, "zip-file", "", "Path to the ZIP file to upload")

	flag.Parse()

	if functionName == "" {
		log.Printf("Invalid argument %s: %s", "function-name", functionName)
		os.Exit(-1)
	}

	if zipFile == "" {
		log.Printf("Invalid argument %s: %s", "zipFile", zipFile)
		os.Exit(-1)
	}

	if _, err := os.Stat(zipFile); errors.Is(err, os.ErrNotExist) {
		log.Printf("Invalid argument %s: %s: File does not exist", "zipFile", zipFile)
		os.Exit(-1)
	}

	awsConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("Failed to load AWS config: %v", err)
		os.Exit(-1)
	}

	bytes, err := os.ReadFile(zipFile)
	if err != nil {
		log.Printf("Failed to read ZIP file: %v", err)
		os.Exit(-1)
	}

	client := lambda.NewFromConfig(awsConfig)

	input := lambda.UpdateFunctionCodeInput{
		FunctionName: aws.String(functionName),
		ZipFile:      bytes,
	}
	_, err = client.UpdateFunctionCode(context.TODO(), &input)
	if err != nil {
		var apiErr smithy.APIError
		if errors.As(err, &apiErr) {
			log.Printf("failed to upload function code (code: %s, message: %s, fault: %s)",
				apiErr.ErrorCode(), apiErr.ErrorMessage(), apiErr.ErrorFault().String(),
			)
			os.Exit(-1)
		}
		log.Printf("failed to upload function code: %v", err)
		os.Exit(-1)
	}
}
