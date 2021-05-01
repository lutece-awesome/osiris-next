package file

import (
	"crypto/sha256"
	"io"
	"os"
)

// SHA256 sum
type Sum256 string

type fileMixin struct {
	// The file name will generated as {current_time_stamp}_{random uint64}
	path string
	// Does this file created
	created bool
}

func getSum256FromIoReader(src io.Reader) (Sum256, error) {
	h := sha256.New()
	_, err := io.Copy(h, src)
	if err != nil {
		return "", err
	}
	return Sum256(h.Sum(nil)), nil
}

func (fileMix *fileMixin) getSum256() (Sum256, error) {
	f, err := os.Open(fileMix.path)
	if err != nil {
		return "", err
	}
	sum256, err := getSum256FromIoReader(f)
	if err != nil {
		return "", err
	}
	err = f.Close()
	if err != nil {
		return sum256, err
	}
	return sum256, nil
}

func (fileMix *fileMixin) emptyFile() (*os.File, error) {
	// Try to remove the former one if exists.
	if _, err := os.Stat(fileMix.path); os.IsExist(err) {
		err = os.Remove(fileMix.path)
		if err != nil {
			return nil, err
		}
	}
	return os.Create(fileMix.path)
}
