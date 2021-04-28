package file

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"sync"
)

type fileManager struct {
	sync.RWMutex
	identifyKeyMap map[string]bool
	workDir        string
}

func (fm *fileManager) checkExist(identifyKey string) bool {
	fm.RLock()
	defer fm.RUnlock()
	var _, ok = fm.identifyKeyMap[identifyKey]
	return ok
}

func (fm *fileManager) DownloadFile(url, identifyKey string) (string, error) {
	var fileName string
	if len(identifyKey) > 0 {
		fileName = identifyKey
		if fm.checkExist(identifyKey) {
			return identifyKey, nil
		}
	} else {
		fileName = fmt.Sprintf("temp_%d", rand.Uint64())
	}
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	} else if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("can not download file because of the http response code is %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	testFile, err := os.Create(path.Join(fm.workDir, fileName))
	if err != nil {
		return "", nil
	}
	defer testFile.Close()
	written, err := io.Copy(testFile, resp.Body)
	if err != nil {
		return "", nil
	}
	log.Printf("successfully download file with %s. [written %d]\n", fileName, written)
	return fileName, nil
}

func NewFileManager(workDir string) (*fileManager, error) {
	fileInfo, err := ioutil.ReadDir(workDir)
	if err != nil {
		return nil, err
	}
	fm := fileManager{
		identifyKeyMap: make(map[string]bool),
		workDir:        workDir,
	}
	for _, file := range fileInfo {
		if !file.IsDir() {
			fm.identifyKeyMap[file.Name()] = true
		}
	}
	log.Printf("cached files with %d\n", len(fm.identifyKeyMap))
	return &fm, nil
}
