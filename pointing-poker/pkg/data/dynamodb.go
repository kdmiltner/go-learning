package data

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type dynamoDatabase struct {
	dynamodb *dynamodb.Client
}

func (d *dynamoDatabase) Write(value any) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("value is not a string")
	}
	fmt.Println(str)
	return nil
}

func (d *dynamoDatabase) Read(value any) (bool, error) {
	str, ok := value.(string)
	if !ok {
		return false, fmt.Errorf("value is not a string")
	}
	fmt.Println(str)
	return false, nil
}
