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

	var bufChars, tmpChars strings.Builder
	var beforeChar rune
	var isBackslash bool

	for _, lineItem := range line {
		if lineItem == '\\' && !isBackslash {
			isBackslash = true
			continue
		}
		if isBackslash {
			bufChars.WriteRune(lineItem)
			beforeChar = lineItem
			isBackslash = false
			continue
		}

		if unicode.IsDigit(lineItem) {
			if beforeChar > 0 {
				tmpChars.WriteRune(lineItem)
			} else {
				return "", ErrInvalidString
			}
		} else {
			numRepeat, err := strconv.Atoi(tmpChars.String())
			if err != nil {
				return "", ErrInvalidString
			}

			if numRepeat > 0 && beforeChar > 0 {
				bufChars.WriteString(strings.Repeat(string(beforeChar), numRepeat-1))
				tmpChars.Reset()
			}
			bufChars.WriteRune(lineItem)
			beforeChar = lineItem
		}
	}

	numRepeat, err := strconv.Atoi(tmpChars.String())
	if err != nil {
		return "", ErrInvalidString
	}
	if numRepeat > 0 && beforeChar > 0 {
		bufChars.WriteString(strings.Repeat(string(beforeChar), numRepeat-1))
	}

	res := bufChars.String()

	if len(res) > 0 {
		return bufChars.String(), nil
	} else {
		return "", ErrInvalidString
	}
}
