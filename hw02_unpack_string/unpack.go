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

	isBackslash := func(lineRune []rune, index *int) bool {
		if lineRune[*index] == '\\' && (unicode.IsDigit(lineRune[*index+1]) || lineRune[*index+1] == '\\') {
			*index++
			return true
		}
		return false
	}

	var bufChars strings.Builder
	lineArray := []rune(line)

	for i := 0; i < len(lineArray); i++ {
		if !isBackslash(lineArray, &i) && unicode.IsDigit(lineArray[i]) {
			return "", ErrInvalidString
		}
		if i < len(lineArray)-1 {
			if unicode.IsDigit(lineArray[i+1]) {
				numRepeat, err := strconv.Atoi(string(lineArray[i+1]))
				if err != nil {
					return "", ErrInvalidString
				}

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
