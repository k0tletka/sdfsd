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
	projectName       = "github.com/k0tletka/sdfsd"
	protobufLocation  = "./internal/protobuf"
	protobufProtoName = "serverapi.proto"
)

var (
	artifactPath = filepath.Join(artifactsBin, artifactName)

	packagesToGenerate = []string{
		"./internal/config",
	}

	buildTools = map[string]string{
		"tools/build_data_generator": "build_data",
	}
)

func Build() error {
	mg.Deps(GenerateGoFiles)
	mg.Deps(GenerateProto)

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

func GenerateGoFiles() error {
	mg.Deps(BuildTools)

	for _, packageToGenerate := range packagesToGenerate {
		if err := sh.RunV("go", "generate", packageToGenerate); err != nil {
			return err
		}
	}

	return nil
}

func GenerateProto() error {
	return sh.RunV(
		"protoc",
		"--go_out="+protobufLocation,
		"--go_opt=paths=source_relative",
		"--go-grpc_out="+protobufLocation,
		"--go-grpc_opt=paths=source_relative",
		filepath.Join(protobufLocation, protobufProtoName),
	)
}

func BuildTools() error {
	for toolLocation, toolName := range buildTools {
		err := sh.RunV("go", "build", "-o",
			filepath.Join(toolLocation, toolName),
			filepath.Join(toolLocation, "main.go"),
		)

		if err != nil {
			return err
		}
	}

	return nil
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
