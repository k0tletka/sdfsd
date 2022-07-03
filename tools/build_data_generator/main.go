package main

import (
	_ "embed"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

//go:embed build_data.template
var buildDataText string

var (
	buildDataTemplate = template.Must(
		template.New("build_data_template").Parse(buildDataText),
	)
)

type Variable struct {
	Name  string
	Value interface{}
}

type TemplateData struct {
	PackageName string
	Data        []Variable
}

func main() {
	data := &TemplateData{
		PackageName: os.Getenv("GOPACKAGE"),
		Data: []Variable{
			{Name: "VersionNumber", Value: getVersionNumber()},
			{Name: "CommitHash", Value: getCommitHash()},
			{Name: "BuildTimeUnix", Value: getBuildTimeUnix()},
		},
	}

	buildDataFileName := os.Getenv("GOFILE")
	buildDataFile, err := os.OpenFile(
		strings.TrimSuffix(buildDataFileName, filepath.Ext(buildDataFileName))+"_generated.go",
		os.O_WRONLY|os.O_CREATE,
		0755,
	)

	if err != nil {
		log.Fatalln(err)
	}

	if err := buildDataTemplate.Execute(buildDataFile, data); err != nil {
		log.Fatalln(err)
	}
}

func getVersionNumber() string {
	tagName, err := exec.Command("git", "describe", "--tags", "--abbrev=0").Output()
	if err != nil {
		log.Fatalln(err)
	}

	return `"` + strings.ReplaceAll(string(tagName), "\n", "") + `"`
}

func getCommitHash() string {
	commitHash, err := exec.Command("git", "rev-parse", "HEAD").Output()
	if err != nil {
		log.Fatalln(err)
	}

	return `"` + strings.ReplaceAll(string(commitHash), "\n", "") + `"`
}

func getBuildTimeUnix() int64 {
	return time.Now().Unix()
}
