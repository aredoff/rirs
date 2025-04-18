package rirs

import (
	"bufio"
	"encoding/json"
	"fmt"

	"github.com/aredoff/rirs/fs"

	"os"
	"path/filepath"

	"github.com/aredoff/rirs/parser"
)

const (
	bufferSize = 5 * 1024 * 1024 // 5MB
)

type storage struct {
	folder  *fs.Folder
	writers map[string]*bufio.Writer
	files   map[string]*os.File
}

func NewStorage(folder *fs.Folder) (*storage, error) {
	storage := &storage{
		folder:  folder,
		writers: make(map[string]*bufio.Writer),
		files:   make(map[string]*os.File),
	}

	// Initialize writers for each type
	types := []string{"asns", "inetnums", "routes", "routes6", "persons", "organizations", "domains"}
	for _, t := range types {
		if err := storage.initWriter(t); err != nil {
			return nil, fmt.Errorf("failed to initialize writer for %s: %w", t, err)
		}
	}

	return storage, nil
}

func (s *storage) initWriter(objType string) error {
	filename := filepath.Join(s.folder.Path(), fmt.Sprintf("%s.json", objType))
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filename, err)
	}

	writer := bufio.NewWriterSize(file, bufferSize)

	if _, err := writer.WriteString("{\n"); err != nil {
		file.Close()
		return fmt.Errorf("failed to write opening bracket: %w", err)
	}

	s.files[objType] = file
	s.writers[objType] = writer
	return nil
}

func (s *storage) SaveASN(asn *parser.ASN) error {
	return s.saveObject("asns", asn.ASNumber, asn)
}

func (s *storage) SaveInetNum(inetnum *parser.InetNum) error {
	return s.saveObject("inetnums", inetnum.IPRange, inetnum)
}

func (s *storage) SaveRoute(route *parser.Route) error {
	return s.saveObject("routes", route.Prefix, route)
}

func (s *storage) SaveRoute6(route6 *parser.Route6) error {
	return s.saveObject("routes6", route6.Prefix, route6)
}

func (s *storage) SavePerson(person *parser.Person) error {
	return s.saveObject("persons", person.NicHdl, person)
}

func (s *storage) SaveOrganization(org *parser.Organization) error {
	return s.saveObject("organizations", org.OrgID, org)
}

func (s *storage) SaveDomain(domain *parser.Domain) error {
	return s.saveObject("domains", domain.Domain, domain)
}

func (s *storage) saveObject(objType, key string, obj interface{}) error {
	writer, ok := s.writers[objType]
	if !ok {
		return fmt.Errorf("writer for %s not initialized", objType)
	}

	jsonData, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("failed to marshal object: %w", err)
	}

	if _, err := writer.WriteString(fmt.Sprintf("  \"%s\": %s", key, string(jsonData))); err != nil {
		return fmt.Errorf("failed to write object: %w", err)
	}

	if writer.Buffered() > bufferSize {
		if err := writer.Flush(); err != nil {
			return fmt.Errorf("failed to flush writer: %w", err)
		}
	}

	return nil
}

func (s *storage) Close() error {
	var lastErr error

	// Close all writers and files
	for objType, writer := range s.writers {
		// Write closing bracket
		if _, err := writer.WriteString("\n}\n"); err != nil {
			lastErr = fmt.Errorf("failed to write closing bracket for %s: %w", objType, err)
			continue
		}

		// Flush and close writer
		if err := writer.Flush(); err != nil {
			lastErr = fmt.Errorf("failed to flush writer for %s: %w", objType, err)
		}

		// Close file
		if file, ok := s.files[objType]; ok {
			if err := file.Close(); err != nil {
				lastErr = fmt.Errorf("failed to close file for %s: %w", objType, err)
			}
		}
	}

	return lastErr
}
