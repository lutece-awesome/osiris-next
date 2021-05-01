package file

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"path"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type FileManager struct {
	// mutex to guarantee safety access cacheMap among multiple goroutines.
	lock sync.Mutex
	// Sum256 to downloadStatus mapping
	cacheMap map[Sum256]*downloadStatus
	// Working directory to store the downloaded files.
	workDir string
	// Random number generator to generate the file name suffix.
	random *rand.Rand
}

func (fm *FileManager) getDownloadStatusOrCreateAndSchedule(sum256 Sum256, url string) *downloadStatus {
	fm.lock.Lock()
	defer fm.lock.Unlock()
	var ds *downloadStatus
	if ret, ok := fm.cacheMap[sum256]; ok && ret.Err == nil {
		ds = ret
	} else {
		ds = newDownloadStatus(fm.generateFileName())
		fm.cacheMap[sum256] = ds
		ds.scheduleDownload(url)
	}
	return ds
}

func (fm *FileManager) generateFileName() string {
	// Generate file name
	return path.Join(fm.workDir, fmt.Sprintf("%v_%v", time.Now().Format(time.RFC3339), fm.random.Uint64()))
}

func NewFileManager(workDir string) (*FileManager, error) {
	fileInfo, err := ioutil.ReadDir(workDir)
	if err != nil {
		return nil, err
	}
	fm := FileManager{
		cacheMap: make(map[Sum256]*downloadStatus),
		workDir:  workDir,
		random:   rand.New(rand.NewSource(time.Now().Unix())),
	}
	log.Info("scan all cached files under ", workDir)
	for _, file := range fileInfo {
		if !file.IsDir() {
			ds := newDownloadStatus(path.Join(workDir, file.Name())).markAsDone()
			sum256, err := ds.FileMixin.getSum256()
			if err != nil {
				return nil, err
			}
			fm.cacheMap[sum256] = ds
		}
	}
	log.Info(fmt.Sprintf("%d cached file loaded", len(fileInfo)))
	return &fm, nil
}
