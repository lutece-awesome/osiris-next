package test

import (
	"testing"

	fi "github.com/lutece-awesome/osiris-next/file"
	util "github.com/lutece-awesome/osiris-next/util"
)

var (
	testDir = "work_test"
)

func prepare() {
	util.MakeDirectoryIfNotExists(testDir)
}

func scheduleDownloadTask(fm *fi.FileManager, url, md5 string, run chan<- error) {
	ds := fm.DownloadFile(url, fi.Md5Sum(md5))
	<-ds.Done
	run <- ds.Err
}

func TestFileDownloadWithoutCache(t *testing.T) {
	prepare()
	fm, err := fi.NewFileManager(testDir)
	if err != nil {
		t.Fatal(err)
	}
	run := make(chan error)
	go scheduleDownloadTask(fm, "http://localhost:9877/static/case1/1.in", "", run)
	go scheduleDownloadTask(fm, "http://localhost:9877/static/case1/2.in", "", run)
	go scheduleDownloadTask(fm, "http://localhost:9877/static/case1/1.out", "", run)
	go scheduleDownloadTask(fm, "http://localhost:9877/static/case1/2.out", "", run)
	for count := 0; count < 4; count++ {
		if err != nil {
			t.Fatal(err)
		}
	}
}
