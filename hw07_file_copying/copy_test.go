package main

import (
	"errors"
	"os"
	"testing"
)

func compareFiles(fileOne, fileTwo string) bool {
	fileOneContent, err := os.ReadFile(fileOne)
	if err != nil {
		return false
	}

	fileTwoContent, err := os.ReadFile(fileTwo)
	if err != nil {
		return false
	}
	return len(fileOneContent) == len(fileTwoContent)
}

func TestCopy(t *testing.T) {
	testCase := []struct {
		name       string
		expOutFile string
		offset     int64
		limit      int64
		expError   error
	}{
		{
			name: "Offset 0 limit 0",

			expOutFile: "testdata/out_offset0_limit0.txt",
			offset:     0,
			limit:      0,
			expError:   nil,
		},
		{
			name:       "Offset 0 limit 10",
			expOutFile: "testdata/out_offset0_limit10.txt",
			offset:     0,
			limit:      10,
			expError:   nil,
		},
		{
			name:       "Offset 0 limit 1000",
			expOutFile: "testdata/out_offset0_limit1000.txt",
			offset:     0,
			limit:      1000,
			expError:   nil,
		},
		{
			name:       "Offset 0 limit 10000",
			expOutFile: "testdata/out_offset0_limit10000.txt",
			offset:     0,
			limit:      10000,
			expError:   nil,
		},
		{
			name:       "Offset 100 limit 1000",
			expOutFile: "testdata/out_offset100_limit1000.txt",
			offset:     100,
			limit:      1000,
			expError:   nil,
		},
		{
			name:       "Offset 6000 limit 1000",
			expOutFile: "testdata/out_offset6000_limit1000.txt",
			offset:     6000,
			limit:      1000,
			expError:   nil,
		},
		{
			name:       "Offset 10000 limit 1000",
			expOutFile: "",
			offset:     10000,
			limit:      1000,
			expError:   ErrOffsetExceedsFileSize,
		},
	}

	inputFile := "testdata/input.txt"

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			tmpFile, err := os.CreateTemp(os.TempDir(), "tmp_")
			if err != nil {
				t.Fatal(err)
			}
			defer func() {
				err = tmpFile.Close()
				if err != nil {
					t.Fatal("Error close file", err)
				}
				err = os.Remove(tmpFile.Name())
				if err != nil {
					t.Fatal("Error remove file", err)
				}
			}()

			err = Copy(inputFile, tmpFile.Name(), tc.offset, tc.limit)

			if tc.expError != nil {
				if !errors.Is(err, tc.expError) {
					t.Errorf("expected error %v, from %v", tc.expError, err)
				}
				return
			}

			if !compareFiles(tmpFile.Name(), tc.expOutFile) {
				if err != nil {
					t.Fatal("Error seek file", err)
				}
				t.Errorf("Files not equal")
			}
		})
	}
}
