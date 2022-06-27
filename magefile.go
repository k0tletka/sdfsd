//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"os"
	"runtime"
)

const (
	artifactsBin = "build"
	artifactName = "sdfsd"
)

func Build() error {
	if enableSystemdUnitFile && runtime.GOOS == "linux" {
		mg.Deps(InstallSystemdUnitFile)
	}

	if enableInitdFile && checkInitdInstalled() {
		mg.Deps(InstallInitdFile)
	}

	return sh.RunV("go", "build", "-o", artifactsBin+artifactName)
}

func Install() error {
	return nil
}

func InstallSystemdUnitFile() error {
	return nil
}

func InstallInitdFile() error {
	return nil
}

func checkInitdInstalled() bool {
	_, err := os.Stat("/sbin/init")
	return !os.IsNotExist(err)
}
