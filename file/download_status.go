package file

// Stored download status within cache map.
type downloadStatus struct {
	// Shared channel of file downloading status, will close if done.
	Done chan struct{}
	// The downloaded file.
	FileMixin fileMixin
	Err       error
}

func (ds *downloadStatus) markAsDone() *downloadStatus {
	close(ds.Done)
	return ds
}

func (ds *downloadStatus) scheduleDownload(url string) {
	go newDownloadTask(url, ds)
}

func newDownloadStatus(willCreatedFile string) *downloadStatus {
	ds := &downloadStatus{
		Done: make(chan struct{}),
		FileMixin: fileMixin{
			path:    willCreatedFile,
			created: false,
		},
	}
	return ds
}
