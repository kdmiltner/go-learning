package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

const (
	// goModPath represents the go.mod path `/go.mod`.
	goModPath = "/go.mod"
	//goModRoot represents the go.mod path at the root directory.
	goModRoot = "go.mod"
	// rootPath sets the root directory path.
	rootPath = "."
	// mainAppName represents the name of the app at the root level.
	mainAppName = "go-learning"
)

// Sets global and env variables that can be accessed in multiple functions.
var (
	appRecords []AppRecord
)

// AppRecord holds the information about an app record. An AppRecord is created for any directory that
// contains a go.mod file.
type AppRecord struct {
	AppPath string
	AppName string
}

func Run() {
	fmt.Println(compileAppRec("magefiles/mage/go.mod"))
}

// HelloCFA is a mage target that prints out "Hello, Chick-fil-A!".
// Run with the command: `mage helloCFA`.
func HelloCFA() {
	fmt.Println("Hello, Chick-fil-A!")
}

// CompileApps will walk the repo directory and compile a list of apps based on the
// go.mod files found. Each go.mod file path will be treated as an app.
// Run with the command: `mage compileApps`.
func CompileApps() error {
	err := compileRecords()
	fmt.Println(appRecords)

	return err
}

// compileRecords walks the root directory to compile a list of `[]AppRecord`.
func compileRecords() error {
	err := filepath.WalkDir(rootPath, func(path string, dir fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Process paths and handle case where path can equal the root go.mod.
		if strings.Contains(path, goModPath) || path == goModRoot {
			record := compileAppRec(path)
			appRecords = append(appRecords, record)
		}

		return nil
	})

	return err
}

// compileAppRec compiles all the app information into an `AppRecord{}`.
func compileAppRec(path string) AppRecord {
	record := AppRecord{}

	// Handle case where the root directory has a go.mod.
	if path == goModRoot {
		record.AppPath = filepath.Base(filepath.Dir(path))
		record.AppName = mainAppName

		return record
	}

	path, _ = strings.CutSuffix(path, goModPath)
	record.AppPath = path
	record.AppName = strings.Replace(path, "/", "-", -1)

	return record
}
