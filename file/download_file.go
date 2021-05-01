package file

func (fm *FileManager) DownloadFile(url string, sum256 Sum256) *downloadStatus {
	var ds *downloadStatus
	if len(sum256) > 0 {
		ds = fm.getDownloadStatusOrCreateAndSchedule(sum256, url)
	} else {
		ds = newDownloadStatus(fm.generateFileName())
		ds.scheduleDownload(url)
	}
	return ds
}
