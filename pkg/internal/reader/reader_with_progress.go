package reader

import "io"

type readerWithProgress struct {
	io.Reader
	bytesRead int
	progressListener func(int)
}

func NewReaderWithProgress(reader io.Reader, progressListener func(int)) *readerWithProgress {
	return &readerWithProgress{
		Reader: reader,
		progressListener: progressListener,
	}
}

func (reader *readerWithProgress) Read(b []byte) (int, error) {
	n, err := reader.Reader.Read(b)
	reader.bytesRead += n
	reader.progressListener(reader.bytesRead)
	return n, err
}