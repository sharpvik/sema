//go:build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

func Install() (err error) {
	version, err := version()
	if err != nil {
		return
	}
	_, err = sh.Output(
		"go", "install",
		"-ldflags", "-w -s -X main.version="+version,
	)
	return
}

func Version() error {
	version, err := version()
	if err != nil {
		return err
	}
	fmt.Println(version)
	return nil
}

func version() (string, error) {
	return sh.Output("git", "describe", "--tags", "--abbrev=0")
}
