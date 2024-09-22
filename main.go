package main

import (
	"fmt"

	"github.com/rifqidaiva/cpm-calculator/internal/pert"
)

func main() {
	var dataInput dataInput

	rawData, err := getFileContent()
	if err != nil {
		panic(err.Error())
	}

	dataInput.data, err = parseRawData(rawData)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(dataInput.getPathList())
	pert.CreatePert(dataInput.getPathList())
}
