package file

import (
	"crypto/md5"
	"io"
	"os"
)

type Md5Sum string

type fileMixin struct {
	// The file name will generated as {current_time_stamp}_{random uint64}
	path string
	// Does this file created
	created bool
}

func getMd5FromIoReader(src io.Reader) (Md5Sum, error) {
	h := md5.New()
	_, err := io.Copy(h, src)
	if err != nil {
		return "", err
	}
	return Md5Sum(h.Sum(nil)), nil
}

func (fileMix *fileMixin) getMd5() (Md5Sum, error) {
	f, err := os.Open(fileMix.path)
	if err != nil {
		return "", err
	}
	md5Val, err := getMd5FromIoReader(f)
	if err != nil {
		return "", err
	}
	err = f.Close()
	if err != nil {
		return md5Val, err
	}
	return md5Val, nil
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
