//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

const (
	projectName = "github.com/k0tletka/sdfsd"
)

var (
	artifactPath = filepath.Join(artifactsBin + artifactName)
)

func Build() error {
	if enableSystemdUnitFile && runtime.GOOS == "linux" {
		mg.Deps(InstallSystemdUnitFile)
	} else if enableInitdFile && checkInitdInstalled() {
		mg.Deps(InstallInitdFile)
	}

	return sh.RunV("go", "build", "-o", artifactPath, projectName)
}

func Install() error {
	if err := os.MkdirAll(installPath, 0755); err != nil {
		return err
	}

	sourceExec, err := ioutil.ReadFile(artifactPath)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filepath.Join(installPath, artifactName), sourceExec, 0755)
}

func InstallSystemdUnitFile() error {
	// TODO: Make installable systemd file
	return nil
}

func InstallInitdFile() error {
	// TODO: Make installable init.d file
	return nil
}

func checkInitdInstalled() bool {
	_, err := os.Stat("/sbin/init")
	return err == nil
}
