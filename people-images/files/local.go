package files

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Local struct {
	basePath    string
	maxFileSize int
}

func NewLocal(basePath string, maxFileSize int) *Local {
	return &Local{
		basePath:    basePath,
		maxFileSize: maxFileSize,
	}
}

func (l *Local) Save(path string, reader io.Reader) error {
	// get the full path
	fullPath := l.fullPath(path)
	fmt.Println("full path")

	// get the directory
	d := filepath.Dir(fullPath)
	err := os.MkdirAll(d, os.ModePerm)
	fmt.Println("mkdir")

	if err != nil {
		return fmt.Errorf("unable to create a directory: %s", err)
	}

	// check whether file exists
	// if file exists, remove it
	_, err = os.Stat(fullPath)
	fmt.Println("os.Stat")

	if err == nil {
		err := os.Remove(fullPath)
		if err != nil {
			return fmt.Errorf("unable to delete a file: %s", err)
		}
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("error get a file info: %s", err)
	}

	// if not, create one
	file, err := os.Create(fullPath)
	fmt.Println("create file")

	if err != nil {
		return fmt.Errorf("unable to create a file: %s", err)
	}
	defer file.Close()

	// copy reader into file
	_, err = io.Copy(file, reader)
	fmt.Println("copy")

	if err != nil {
		return fmt.Errorf("error writing into a file: %s", err)
	}

	return nil
}

func (l *Local) fullPath(path string) string {
	return filepath.Join(l.basePath, path)
}
