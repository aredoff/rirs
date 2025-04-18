package parser

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	bufferSize = 5 * 1024 * 1024 // 5MB
)

func (p *Parser) ParseFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	return p.parseFromReader(file)
}

func (p *Parser) ParseGZFile(gzFilePath string) error {
	gzFile, err := os.Open(gzFilePath)
	if err != nil {
		return fmt.Errorf("failed to open gzip file %s: %w", gzFilePath, err)
	}
	defer gzFile.Close()

	gzReader, err := gzip.NewReader(gzFile)
	if err != nil {
		return fmt.Errorf("failed to create gzip reader for %s: %w", gzFilePath, err)
	}
	defer gzReader.Close()

	return p.parseFromReader(gzReader)
}

func (p *Parser) parseFromReader(reader io.Reader) error {
	var currentObject []string
	var currentType string

	scanner := bufio.NewScanner(reader)
	buf := make([]byte, 0, bufferSize)
	scanner.Buffer(buf, bufferSize)

	for scanner.Scan() {
		line := scanner.Text()

		// Skip empty lines at the start
		if len(currentObject) == 0 && line == "" {
			continue
		}

		// If we hit an empty line and have content, parse the current object
		if line == "" && len(currentObject) > 0 {
			if err := p.parseAndSaveObject(currentType, currentObject); err != nil {
				return fmt.Errorf("failed to parse object: %w", err)
			}
			currentObject = nil
			currentType = ""
			continue
		}

		// If this is the first line of a new object, determine its type
		if len(currentObject) == 0 {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				currentType = strings.TrimSpace(parts[0])
			}
		}

		currentObject = append(currentObject, line)
	}

	// Parse the last object if there is one
	if len(currentObject) > 0 {
		if err := p.parseAndSaveObject(currentType, currentObject); err != nil {
			return fmt.Errorf("failed to parse object: %w", err)
		}
	}
	return scanner.Err()
}
