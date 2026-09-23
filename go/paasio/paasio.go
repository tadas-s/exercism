package paasio

import (
	"io"
	"sync"
)

type writeCounter struct {
	writer io.Writer

	writes   int64
	writeOps int

	mutex *sync.RWMutex
}

type readCounter struct {
	reader io.Reader

	reads   int64
	readOps int

	mutex *sync.RWMutex
}

type readWriteCounter struct {
	ReadCounter
	WriteCounter
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{writer, 0, 0, new(sync.RWMutex)}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader, 0, 0, new(sync.RWMutex)}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		NewReadCounter(readwriter),
		NewWriteCounter(readwriter),
	}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	rc.mutex.Lock()
	defer rc.mutex.Unlock()

	read, err := rc.reader.Read(p)

	rc.readOps += 1
	rc.reads += int64(read)

	return read, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.mutex.Lock()
	defer rc.mutex.Unlock()

	return rc.reads, rc.readOps
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	wc.mutex.Lock()
	defer wc.mutex.Unlock()

	written, err := wc.writer.Write(p)

	wc.writeOps += 1
	wc.writes += int64(written)

	return written, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	wc.mutex.Lock()
	defer wc.mutex.Unlock()

	return wc.writes, wc.writeOps
}
