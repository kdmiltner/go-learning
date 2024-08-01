package data

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type dynamoDatabase struct {
	dynamodb *dynamodb.Client
}

func (d *dynamoDatabase) Write(str string) error {
	fmt.Println(str)
	return nil
}

func (d *dynamoDatabase) Read(str string) (bool, error) {
	fmt.Println(str)
	return false, nil
}
