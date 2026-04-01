package data

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"sync"
)

const (
	tmpDBFile = "tmp_database.csv"
)

type csvDatabase struct {
	csv *os.File
}

func (c *csvDatabase) Write(ctx context.Context, value string) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("context error: %v", err)
	}

	_, err := c.csv.Write([]byte(fmt.Sprintf("%s,\n", value)))
	if err != nil {
		return fmt.Errorf("error writing to %s: %v", tmpDBFile, err)
	}

	return nil
}

func (c *csvDatabase) Read(ctx context.Context, value string) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, fmt.Errorf("context error: %v", err)
	}

	var (
		found bool
		wg    sync.WaitGroup
	)

	fmt.Println(c.csv.Name())
	r := csv.NewReader(c.csv)
	records, err := r.ReadAll()
	if err != nil {
		return found, fmt.Errorf("error reading from %s: %v", tmpDBFile, err)
	}
	if records == nil {
		return found, errors.New("records is empty")
	}

	// TODO: look at using the index to capture where the sessionID match is
	for _, rec := range records {
		wg.Add(1)
		go func() bool {
			defer wg.Done()
			if rec[0] == value {
				found = true
			}
			return found
		}()
		if found == true {
			break
		}
	}
	wg.Wait()

	return found, nil
}
