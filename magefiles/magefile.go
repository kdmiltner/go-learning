package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/magefile/mage/sh"
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

// CompileApps will walk the repo directory and compiles a list of apps based on the
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

// RunPython with run the defined python script.
// Run with the command: `mage runPython`
func RunPython() error {
	err := sh.RunV("python", "./apex/interfaces/1_python/1_python_animals.py")

	return err
}
