//go:build mage

// Use this file to produce relase binaries for `sema`.
// @ https://github.com/sharpvik/sema
//
// It relies on a build of the `mage` tool.
// @ https://github.com/magefile/mage
//
// Once you have `mage` command installed, you can run:
// $ mage bins
//
// This will create a folder called `bin` and put all binary executables there.

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/sh"
)

var variants = [5]struct {
	os   string
	arch string
}{
	{"darwin", "amd64"},
	{"linux", "amd64"},
	{"windows", "amd64"},
	{"linux", "386"},
	{"windows", "386"},
}

func extension(os string) string {
	if os == "windows" {
		return ".exe"
	}
	return ""
}

func exectableName(os, arch string) string {
	return fmt.Sprintf("./bin/%s_%s_sema%s", os, arch, extension(os))
}

func envMap(os, arch string) map[string]string {
	return map[string]string{
		"GOOS":   os,
		"GOARCH": arch,
	}
}

func crossCompileAll() (err error) {
	for _, v := range variants {
		_, err := sh.Exec(envMap(v.os, v.arch), os.Stdout, os.Stdout,
			"go", "build", "-o", exectableName(v.os, v.arch), ".")
		if err != nil {
			return err
		}
	}
	return
}

func makeBinDir() (err error) {
	return os.MkdirAll("bin", 0777)
}

func Bins() (err error) {
	if err = makeBinDir(); err != nil {
		return
	}
	return crossCompileAll()
}
