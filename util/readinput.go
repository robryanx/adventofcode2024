package util

import (
	"bytes"
	"fmt"
	"iter"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

type Iterator = func(s []byte)

func filename(dayPart string, isSample bool) string {
	folder := "inputs"
	if isSample {
		folder = "samples"
	}

	_, fileName, _, _ := runtime.Caller(0)
	prefixPath := filepath.Dir(fileName)

	return fmt.Sprintf("%s/../%s/%s.txt", prefixPath, folder, dayPart)
}

func ReadBytes(dayPart string, isSample bool) ([]byte, error) {
	return os.ReadFile(filename(dayPart, isSample))
}

func ReadStrings(dayPart string, isSample bool, delim string) (iter.Seq[string], error) {
	partIter, err := read(filename(dayPart, isSample), delim)
	if err != nil {
		return nil, err
	}

	return func(yield func(string) bool) {
		for part := range partIter {
			partStr := string(part)
			if !yield(partStr) {
				return
			}
		}
	}, nil
}

func ReadInts(dayPart string, isSample bool, delim string) (iter.Seq[int], error) {
	partIter, err := read(filename(dayPart, isSample), delim)
	if err != nil {
		return nil, err
	}

	return func(yield func(int) bool) {
		for part := range partIter {
			partInt, err := strconv.Atoi(string(part))
			if err != nil {
				return
			}
			if !yield(partInt) {
				return
			}
		}
	}, nil
}

func read(file string, delim string) (iter.Seq[[]byte], error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return func(yield func([]byte) bool) {
		for _, row := range bytes.Split(b, []byte(delim)) {
			if !yield(row) {
				return
			}
		}
	}, nil
}
