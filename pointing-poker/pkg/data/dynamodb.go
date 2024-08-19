package data

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type dynamoDatabase struct {
	dynamodb *dynamodb.Client
}

func (d *dynamoDatabase) Write(value string) error {
	fmt.Println(value)
	return nil
}

func (d *dynamoDatabase) Read(value string) (bool, error) {
	fmt.Println(value)
	return false, nil
}
