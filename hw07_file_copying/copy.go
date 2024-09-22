package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	inputFile, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer inputFile.Close()

	statFile, err := inputFile.Stat()
	if err != nil {
		return ErrUnsupportedFile
	}

	fileSize := statFile.Size()
	if limit != 0 {
		if limit+offset >= fileSize {
			fileSize -= offset
		} else {
			fileSize = limit
		}
	}
	if fileSize < 0 {
		return ErrOffsetExceedsFileSize
	}

	_, err = inputFile.Seek(offset, 0)
	if err != nil {
		return ErrOffsetExceedsFileSize
	}

	finalFile, err := os.OpenFile(toPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.FileMode(0o666))
	if err != nil {
		return fmt.Errorf("%w: %w", ErrUnsupportedFile, err)
	}
	defer finalFile.Close()

	err = writeProgress(inputFile, finalFile, fileSize)
	if err != nil {
		return err
	}

	return nil
}

func writeProgress(inputFile *os.File, finalFile *os.File, fileSize int64) error {
	var progressBar *pb.ProgressBar
	ioWriter := io.Writer(finalFile)
	ioReader := io.Reader(inputFile)

	progressBar = pb.StartNew(int(fileSize))
	ioReader = progressBar.NewProxyReader(inputFile)

	_, err := io.CopyN(ioWriter, ioReader, fileSize)
	if err != nil {
		return err
	}

	progressBar.Finish()

	return nil
}
