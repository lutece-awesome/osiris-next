package file

func (fm *FileManager) DownloadFile(url string, md5 Md5Sum) *downloadStatus {
	var ds *downloadStatus
	if len(md5) > 0 {
		ds = fm.getDownloadStatusOrCreateAndSchedule(md5, url)
	} else {
		ds = newDownloadStatus(fm.generateFileName())
		ds.scheduleDownload(url)
	}
	return ds
}
