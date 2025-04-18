package rirs

import (
	"github.com/aredoff/rirs/parser"
)

func (r *rir) Sync() error {
	for _, source := range sources {
		downloadDir, err := r.downloadFolder.SubFolder(source.Name)
		if err != nil {
			return err
		}
		databaseDir, err := r.databaseFolder.SubFolder(source.Name)
		if err != nil {
			return err
		}
		storage, err := NewStorage(databaseDir)
		if err != nil {
			return err
		}
		defer storage.Close()

		parser := parser.NewParser(storage)
		for _, url := range source.httpDatabases {
			filePath, err := downloadFile(downloadDir.Path(), url)
			if err != nil {
				return err
			}
			err = parser.ParseGZFile(filePath)
			if err != nil {
				return err
			}
		}
		err = downloadDir.Clear()
		if err != nil {
			return err
		}
	}
	return nil
}
