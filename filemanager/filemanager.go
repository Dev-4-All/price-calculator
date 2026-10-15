package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	inputFilePath  string
	outputFilePath string
}

func New(inputFilePath, outputFilePath string) *FileManager {
	return &FileManager{
		inputFilePath,
		outputFilePath,
	}
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.inputFilePath)

	if err != nil {
		return nil, errors.New("Failed to open file.")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		return nil, errors.New("Failed to read file content.")
	}

	return lines, nil
}

func (fm FileManager) WriteResults(data any) error {
	file, err := os.Create(fm.outputFilePath)

	if err != nil {
		return errors.New("Failed to create file.")
	}

	defer file.Close()

	time.Sleep(3 * time.Second) // Simulate a time-taking task

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("Failed to convert data into JSON.")
	}

	return nil
}
