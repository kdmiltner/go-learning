package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const (
	// goMod represents any go.mod file.
	goMod = "go.mod"
	// rootPath sets the root directory path.
	rootPath = "."
	// mainAppName represents the name of the app at the root level.
	mainAppName = "go-learning"
)

// compileRecords walks the root directory to compile a list of `[]AppRecord`.
func compileRecords() error {
	clearExistingRecords()
	err := filepath.WalkDir(rootPath, func(path string, dir fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Process paths and handle case where path can equal the root go.mod.
		if strings.Contains(path, goMod) {
			record := compileAppRec(path)
			appRecords = append(appRecords, record)
		}

		return nil
	})

	return err
}

// clearExistingRecords will reset the appRecords slice to the zero value.
func clearExistingRecords() {
	appRecords = []AppRecord{}
}

// compileAppRec compiles all the app information into an `AppRecord{}`.
func compileAppRec(path string) AppRecord {
	record := AppRecord{}

	// Handle case where the root directory includes a go.mod.
	if path == goMod {
		record.AppPath = filepath.Base(filepath.Dir(path))
		record.AppName = mainAppName

		return record
	}

	path, _ = strings.CutSuffix(path, "/"+goMod)
	record.AppPath = path
	record.AppName = strings.Replace(path, "/", "-", -1)

	return record
}

// testAll loops through all appRecords and calls the Test target with an appName.
func testAll() {
	for _, appRec := range appRecords {
		mg.SerialDeps(mg.F(Test, appRec.AppName))
	}
}

// test takes an AppRecord and will test that specific app.
func test(app AppRecord) error {
	const unableToTestErrPrefix = "unable to test"
	rootDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("%s: %s - error: %v", unableToTestErrPrefix, app.AppName, err)
	}

	if err := os.Chdir(app.AppPath); err != nil {
		return fmt.Errorf("%s: %s - error: %v", unableToTestErrPrefix, app.AppName, err)
	}

	fmt.Printf("========= testing %s =========\n", app.AppName)
	if err := sh.RunWithV(cgoEnabled, "gotestsum", "--junitfile", "unit-tests.xml", "--", "-coverprofile=pr.out", "./..."); err != nil {
		return fmt.Errorf("%s: %s - error: %v", unableToTestErrPrefix, app.AppName, err)
	}

	if err := os.Chdir(rootDir); err != nil {
		return err
	}

	return nil
}
