package data

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDatabase(t *testing.T) {
	testCases := []struct {
		testName string
		database string

		expected    ReadWriter
		expectedErr error
	}{
		{
			testName: "Success - csv",
			database: DatabaseCSV,

			expected:    &csvDatabase{},
			expectedErr: nil,
		},
		{
			testName: "Success - dynamodb",
			database: DatabaseDynamoDB,

			expected:    &dynamoDatabase{},
			expectedErr: nil,
		},
		{
			testName: "Failure - unknown database",
			database: "test",

			expected:    nil,
			expectedErr: errors.New("unknown database type: test"),
		},
	}

	for _, test := range testCases {
		t.Run(test.testName, func(t *testing.T) {
			rw, err := NewDatabase(test.database)

			assert.IsType(t, test.expected, rw)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}
