package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	appName = "GrepGrep"
	version = "0.1.0"
	reset = "\033[0m"
	danger = "\033[31m"
	info = "\033[36m"
)

func main() {
	var versionInput bool
	const versionUsage = "Show the version."
	flag.BoolVar(&versionInput, "version", false, versionUsage)
	flag.BoolVar(&versionInput, "v", false, versionUsage)

	flag.Parse()

	if versionInput == true {
		fmt.Println(info + appName + " " + version + reset)
		os.Exit(0)
	}

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println(danger + " You must supply 2 arguments." + reset)
		os.Exit(0)
	}

	fileList := make([]string, 0)
	err = filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)

		return err
	})

	if err != nil {
		panic(err)
	}

	var hasResults = false
	for _, file := range fileList {
		var argumentResults [] string

		for _, argument := range args {
			if fileContains(file, argument) {
				argumentResults = append(argumentResults, file)
			}
		}

		if len(argumentResults) == len(args) {
			fmt.Println(info + file + reset)
			hasResults = true
		}
	}

	if !hasResults {
		fmt.Println(danger + " No results found." + reset)
	}
}

func fileContains(file, text string) bool {
	stat, err := os.Stat(file)
	if err != nil {
		panic(err)
	}

	mode := stat.Mode()
	if mode.IsDir() {
		return false
	}

	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	s := string(fileBytes)

	return strings.Contains(s, text)
}
