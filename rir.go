package rirs

import "github.com/aredoff/rirs/fs"

func New(folder *fs.Folder) (*rir, error) {
	downloadFolder, err := folder.SubFolder("download")
	if err != nil {
		return nil, err
	}
	extractFolder, err := folder.SubFolder("extract")
	if err != nil {
		return nil, err
	}
	databaseFolder, err := folder.SubFolder("database")
	if err != nil {
		return nil, err
	}

	return &rir{
		downloadFolder: downloadFolder,
		extractFolder:  extractFolder,
		databaseFolder: databaseFolder,
	}, nil
}

type rir struct {
	downloadFolder *fs.Folder
	extractFolder  *fs.Folder
	databaseFolder *fs.Folder
}
