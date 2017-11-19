package yml_operator

import (
	"strings"
	"errors"
)

func insertSpaces(slices []string, spacesCount int) []string {
	toInsert := strings.Repeat(" ", spacesCount)
	slicesToReturn := make([]string, 0)
	for _, el := range slices {
		slicesToReturn = append(slicesToReturn, toInsert+el)
	}
	return slicesToReturn
}

func AppendToYmlInSection(toAppend string, sourceYml string, appendPath string) (string, error) {
	splitPath := strings.Split(appendPath, ".")
	splitInput := strings.Split(sourceYml, "\n")
	splitAppend := strings.Split(toAppend, "\n")

	foundIndex, spacesCount, firstSpacesCount := findSpacesCount(splitPath, splitInput)
	if foundIndex >= 0 {
		if firstSpacesCount == -1 {
			firstSpacesCount = 3
		}
		toMerge := insertSpaces(splitAppend, spacesCount+firstSpacesCount)
		merged := append(splitInput[:foundIndex+1], append(toMerge, splitInput[foundIndex+1:]...)...)
		toReturn := ""
		for _, el := range merged {
			toReturn += el + "\n"
		}
		return toReturn, nil
	} else {
		errors.New("Path not found in input yml")
	}
	return "", nil
}
func findSpacesCount(splitPath []string, splitInput []string) (int, int, int) {
	currentSpacesCount := 0
	foundIndex := -1
	firstSpaceCount := -1
	for index, pathElement := range splitPath {
		for _, inputElement := range splitInput {
			spaces := strings.Count(inputElement, " ")
			if spaces > 0 {
				if firstSpaceCount == -1 {
					firstSpaceCount = spaces
				}
			}
			if strings.Contains(inputElement, pathElement) {
				spaceCount := strings.Count(inputElement, " ")
				if spaceCount < currentSpacesCount {
					break
				} else {
					currentSpacesCount = spaceCount
				}
				if index == (len(splitPath) - 1) {
					foundIndex = index
				}
			}
		}
	}

	return foundIndex, currentSpacesCount, firstSpaceCount
}
