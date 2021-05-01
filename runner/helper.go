package runner

import (
	constant "github.com/lutece-awesome/osiris-next/constant"
	fm "github.com/lutece-awesome/osiris-next/file"
	log "github.com/sirupsen/logrus"
)

var FileManager, err = fm.NewFileManager(*constant.Work)

func init() {
	if err != nil {
		log.Error("Unexpected error happened while bootstrap the file manager")
		panic(err)
	}
}
