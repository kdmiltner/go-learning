package data

import (
	"encoding/csv"
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

func (c *csvDatabase) Write(value any) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("value is not a string")
	}
	_, err := c.csv.Write([]byte(fmt.Sprintf("%s,", str)))
	if err != nil {
		return fmt.Errorf("error writing to %s: %v", tmpDBFile, err)
	}

	return nil
}

func (c *csvDatabase) Read(value any) (bool, error) {
	var (
		found bool
		wg    sync.WaitGroup
	)

	str, ok := value.(string)
	if !ok {
		return found, fmt.Errorf("value is not a string")
	}

	r := csv.NewReader(c.csv)
	records, err := r.ReadAll()
	if err != nil {
		return found, fmt.Errorf("error reading from %s: %v", tmpDBFile, err)
	}

	// TODO: look at using the index to capture where the sessionID match is
	for _, rec := range records[0] {
		wg.Add(1)
		go func() bool {
			defer wg.Done()
			if rec == str {
				found = true
			}

			return found
		}()
		wg.Wait()
		if found == true {
			break
		}
	}

	return found, nil
}
