package chunkserver

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
)

// File represents an open file
type File struct {
	FilePointer *os.File
}

// FileHandler struct maintains handles file operations
type FileHandler struct {
	OpenFiles map[string]*File
	lck       sync.Mutex
}

// Open a file
func (fh *FileHandler) Open(filename string) error {
	fh.lck.Lock()
	defer fh.lck.Unlock()

	if _, ok := fh.OpenFiles[filename]; ok {
		return fmt.Errorf("File already open")
	}

	fp, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	fh.OpenFiles[filename] = &File{
		FilePointer: fp,
	}
	return nil
}

// Close a file
func (fh *FileHandler) Close(filename string) error {
	file, err := GetOpenFileSafe(fh, filename, true)
	if err != nil {
		return err
	}

	if err := file.FilePointer.Close(); err != nil {
		return err
	}

	return nil
}

// Read from a file
func (fh *FileHandler) Read(filename string, bytes, offset int64) ([]byte, error) {
	file, err := GetOpenFileSafe(fh, filename, false)
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, bytes)

	// Don't need to lock when doing read/write operations since master server will not allow
	// another process to write and read at the same time
	n, err := file.FilePointer.ReadAt(buffer, offset)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, fmt.Errorf("FileHandler: error reading file %s with %v", filename, err)
	}
	fmt.Println(buffer)
	return buffer[:n], nil
}

// Append data a file
func (fh *FileHandler) Append(filename string, buffer []byte) error {
	file, err := GetOpenFileSafe(fh, filename, false)
	if err != nil {
		return err
	}

	// Don't need to lock when doing read/write operations since master server will not allow
	// another process to write and read at the same time
	if _, err := file.FilePointer.Write(buffer); err != nil && !errors.Is(err, io.EOF) {
		return fmt.Errorf("FileHandler: error write to file %s with %v", filename, err)
	}

	return nil
}

// GetOpenFileSafe executes a function so that it's sychronized
func GetOpenFileSafe(filehandler *FileHandler, filename string, remove bool) (*File, error) {
	filehandler.lck.Lock()
	defer filehandler.lck.Unlock()
	file, ok := filehandler.OpenFiles[filename]
	if !ok {
		return nil, fmt.Errorf("FileHandler: File %s has not been opened", filename)
	}

	if remove {
		delete(filehandler.OpenFiles, filename)
	}
	return file, nil
}
