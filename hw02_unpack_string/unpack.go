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

	var bufChars strings.Builder
	var beforeChar rune

	for key, lineItem := range line {
		if key == 0 && unicode.IsDigit(lineItem) {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(beforeChar) && unicode.IsDigit(lineItem) {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(lineItem) {
			numRepeat, _ := strconv.Atoi(string(lineItem))
			if numRepeat != 0 {
				repeatLine := strings.Repeat(string(beforeChar), numRepeat-1)
				bufChars.WriteString(repeatLine)
			} else {
				tmpStr := bufChars.String()
				tmpStr = tmpStr[:len(tmpStr)-1]
				bufChars.Reset()
				bufChars.WriteString(tmpStr)
			}
		} else {
			bufChars.WriteRune(lineItem)
		}

		beforeChar = lineItem
	}
	return bufChars.String(), nil
}
