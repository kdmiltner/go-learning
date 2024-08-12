package main

import (
	"errors"
	"fmt"
	"log"
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
	cgoEnabled map[string]string
)

// AppRecord holds the information about an app record. An AppRecord is created for any directory that
// contains a go.mod file.
type AppRecord struct {
	AppPath string
	AppName string
}

func init() {
	cgoEnabled = map[string]string{"CGO_ENABLED": "0"}
	if err := CompileApps(); err != nil {
		log.Fatal(err)
	}
}

func Run() {
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

// Test will run tests for a single app or all apps based on the value provided.
// Run with the command: `mage all` OR `mage appName`.
func Test(app string) error {
	if app == "all" {
		testAll()
	}

	var wrapErr error
	for _, appRec := range appRecords {
		if app == appRec.AppPath || app == appRec.AppName {
			if err := test(appRec); err != nil {
				wrapErr = errors.Join(err, wrapErr)
			}
		}
	}

	return wrapErr
}
