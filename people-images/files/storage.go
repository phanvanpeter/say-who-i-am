package files

import "io"

type Storage interface {
	Save(path string, reader io.Reader) error
}
