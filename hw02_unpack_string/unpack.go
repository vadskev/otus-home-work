package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(line string) (string, error) {
	if _, err := strconv.Atoi(line); err == nil {
		return "", ErrInvalidString
	}

	if len(line) == 0 {
		return "", nil
	}

	var bufChars strings.Builder
	lineArray := []rune(line)

	for i := 0; i < len(lineArray); i++ {
		if unicode.IsDigit(lineArray[i]) {
			if (unicode.IsDigit(lineArray[i+1]) || lineArray[i+1] == '\\') && lineArray[i] == '\\' {
				i++
			} else {
				return "", ErrInvalidString
			}
		}

		if i < len(lineArray)-1 {
			if unicode.IsDigit(lineArray[i+1]) {
				numRepeat, _ := strconv.Atoi(string(lineArray[i+1]))
				for j := 0; j < numRepeat; j++ {
					bufChars.WriteRune(lineArray[i])
				}
				i++
				continue
			}
		}
		bufChars.WriteRune(lineArray[i])
	}
	return bufChars.String(), nil
}
