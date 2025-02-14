package slices

import (
	"errors"
)

func GetLastElementInSlice(slice []string) (string, error) {
	sliceLen := len(slice)
	if sliceLen == 0 {
		return "", errors.New("slice is empty")
	}
	return slice[sliceLen-1], nil
}

func GetFirstElementInSlice(slice []string) (string, error) {
	sliceLen := len(slice)
	if sliceLen == 0 {
		return "", errors.New("slice is empty")
	}
	return slice[0], nil
}

func GetSubsetOfSlice(slice []string, startIdx int, count int) (newSlice []string, err error) {
	sliceLen := len(slice)
	if sliceLen == 0 {
		return nil, errors.New("slice is empty")
	}
	if startIdx > sliceLen {
		return nil, errors.New("start index is greater than length")
	}
	end := startIdx + count
	if end > sliceLen {
		return nil, errors.New("end index is greater than length")
	}
	newSlice = slice[startIdx:end]
	return newSlice, nil
}

func GetSliceSplitInHalf(slice []string) (firstHalf []string, secondHalf []string, err error) {
	sliceLen := len(slice)
	a := 0
	if sliceLen%2 != 0 {
		a = 1
	}
	if sliceLen == 0 {
		return nil, nil, errors.New("slice is empty")
	}
	firstHalf = slice[:sliceLen/2+a]
	secondHalf = slice[sliceLen/2+a:]
	err = nil
	return
}

func ChuckSlice(slice []string, chunkSize int) (Chucked [][]string, err error) {
	if len(slice) == 0 {
		return nil, errors.New("slice is empty")
	}
	if chunkSize <= 0 {
		return nil, errors.New("chuck size has to be positive")
	}
	divided := make([][]string, (len(slice)+chunkSize-1)/chunkSize)
	currentIndex := 0
	i := 0
	till := len(slice) - chunkSize
	for currentIndex < till {
		endIndex := currentIndex + chunkSize
		divided[i] = slice[currentIndex:endIndex]
		currentIndex = endIndex
		i++
	}
	divided[i] = slice[currentIndex:]
	return divided, nil
}

func AppendSliceToSlice(slice []string, appendValues []string) []string {
	for _, value := range appendValues {
		slice = append(slice, value)
	}
	return slice
}

func AppendVariadicValueToSlice(slice []string, appendValues ...string) []string {
	appendedSlice := append(slice, appendValues...)
	return appendedSlice
}

func AppendValueToSlice(slice []string, appendValue string) []string {

	appendedSlice := append(slice, appendValue)
	return appendedSlice
}
