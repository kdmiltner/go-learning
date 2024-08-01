package data

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

const (
	DatabaseCSV      = "csv"
	DatabaseDynamoDB = "dynamodb"
)

type ReadWriter interface {
	Write(string) error
	Read(string) (bool, error)
}

func NewDatabase(databaseType string) (ReadWriter, error) {
	switch databaseType {
	case DatabaseCSV:
		var err error
		csv, err := os.OpenFile(tmpDBFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		// TODO: defer csv.Close() where should it go?
		if err != nil {
			return nil, err
		}

		return &csvDatabase{csv}, nil
	case DatabaseDynamoDB:
		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		dynamoDBClient := dynamodb.NewFromConfig(cfg)

		return &dynamoDatabase{dynamoDBClient}, nil
	default:
		return nil, fmt.Errorf("unknown database type: %s", databaseType)
	}
}
