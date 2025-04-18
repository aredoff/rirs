package fs

import (
	"io"
	"os"
	"path/filepath"
	"time"
)

func New(path string) (*Folder, error) {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return nil, err
	}
	return &Folder{path: path}, nil
}

type Folder struct {
	path string
}

func (f *Folder) Path() string {
	return f.path
}

func (f *Folder) Remove() error {
	return os.RemoveAll(f.path)
}

func (f *Folder) Clear() error {
	files, err := os.ReadDir(f.path)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			err = os.RemoveAll(filepath.Join(f.path, file.Name()))
			if err != nil {
				return err
			}
		} else {
			err = os.Remove(filepath.Join(f.path, file.Name()))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (f *Folder) Exist(name string) bool {
	_, err := os.Stat(f.GetPath(name))
	return err == nil
}

func (f *Folder) LastModified(name string) (time.Time, error) {
	info, err := os.Stat(f.GetPath(name))
	if err != nil {
		return time.Time{}, err
	}
	return info.ModTime(), nil
}

func (f *Folder) SubFolder(name string) (*Folder, error) {
	err := os.MkdirAll(f.GetPath(name), 0755)
	if err != nil {
		return nil, err
	}

	return &Folder{
		path: filepath.Join(f.path, name),
	}, nil
}

func (f *Folder) GetPath(name string) string {
	return filepath.Join(f.path, name)
}

func (f *Folder) GetContent(name string) ([]byte, error) {
	file, err := os.Open(f.GetPath(name))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return io.ReadAll(file)
}

func (f *Folder) GetOsFile(name string) (*os.File, error) {
	return os.Open(f.GetPath(name))
}

func (f *Folder) PutContent(name string, content []byte) error {
	return os.WriteFile(f.GetPath(name), content, 0644)
}

func (f *Folder) PutContentAsReader(name string, content io.Reader) error {
	file, err := os.Create(f.GetPath(name))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, content)
	return err
}
