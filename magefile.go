//go:build mage

package main

import (
	"github.com/magefile/mage/sh"
)

func Install() error {
	version, err := sh.Output("git", "describe", "--tags", "--abbrev=0")
	if err != nil {
		return err
	}
	_, err = sh.Output("go", "install", "-ldflags", "-X main.version="+version)
	if err != nil {
		return err
	}
	return nil
}
