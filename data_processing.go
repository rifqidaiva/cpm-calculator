package main

import (
	"fmt"
	"strings"
)

type dataInput struct {
	data [][3]string
}

func (d dataInput) getPredecessors(activity string) []string {
	predecessors := []string{}

	for _, row := range d.data {
		if row[0] == activity {
			data := strings.Split(row[1], ",")
			if data[0] == "-" {
				return []string{}
			}

			predecessors = append(predecessors, data...)
			break
		}
	}

	return predecessors
}

func (d dataInput) getPathList() [][]string {
	pathList := [][]string{}
	endActivity := d.data[len(d.data)-1][0]
	paths := d.findPaths(endActivity, []string{})

	for _, path := range paths {
		reverse(path)
		pathList = append(pathList, path)
	}

	return pathList
}

// Recursive function to find paths.
func (d dataInput) findPaths(activity string, currentPath []string) [][]string {
	currentPath = append(currentPath, activity)
	predecessors := d.getPredecessors(activity)

	if len(predecessors) == 0 {
		return [][]string{currentPath}
	}

	allPaths := [][]string{}

	for _, predecessor := range predecessors {
		paths := d.findPaths(predecessor, currentPath)
		allPaths = append(allPaths, paths...)
	}

	return allPaths
}

func reverse(arr []string) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}

// TODO: Write more data input validation.
func parseRawData(rawData string) ([][3]string, error) {
	parsedData := [][3]string{}
	lines := strings.Split(rawData, "\n")

	for index, line := range lines {
		cleanLine := strings.ReplaceAll(line, " ", "")
		parts := strings.Split(cleanLine, ";")

		if len(parts) != 3 {
			return nil, fmt.Errorf("wrong format at line: %d\nvalue: %s/n", index+1, line)
		}
		parsedData = append(parsedData, [3]string{parts[0], parts[1], parts[2]})
	}

	return parsedData, nil
}
