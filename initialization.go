package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func getArgs() (string, error) {
	flag.Usage = func() {
		flag.PrintDefaults()
	}

	flag.Parse()
	args := flag.Args()

	if len(args) > 1 {
		return "", errors.New("too many arguments")
	}

	if len(args) == 0 {
		fmt.Println("cpm-calculatorV1.0")
		os.Exit(0)
	}

	return args[0], nil
}

func getWorkingDirectory() (string, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return workingDir, nil
}

func getFilePath() (string, error) {
	workingDirectory, err := getWorkingDirectory()
	if err != nil {
		return "", err
	}

	args, err := getArgs()
	if err != nil {
		return "", err
	}

	path := workingDirectory + "\\" + args
	return path, nil
}

func getFileContent() (string, error) {
	filePath, err := getFilePath()
	if err != nil {
		return "", err
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
