package file

import (
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func newDownloadTask(url string, ds *downloadStatus) {
	defer ds.markAsDone()
	if len(ds.FileMixin.path) == 0 {
		ds.Err = fmt.Errorf("empty file path of download call from %v", ds)
		return
	}
	log.Info(fmt.Sprintf("http request of %v scheduled", url))
	resp, err := http.Get(url)
	if err != nil {
		ds.Err = err
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		ds.Err = fmt.Errorf("download http response code error, expected %v, but %v", http.StatusOK, resp.StatusCode)
		return
	}
	testFile, err := ds.FileMixin.emptyFile()
	if err != nil {
		ds.Err = err
		return
	}
	defer testFile.Close()
	_, err = io.Copy(testFile, resp.Body)
	if err != nil {
		ds.Err = err
		return
	}
	log.Info(fmt.Sprintf("http request of %v done", url))
	ds.FileMixin.created = true
	ds.Err = nil
}
