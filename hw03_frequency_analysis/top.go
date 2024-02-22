package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type Words struct {
	word      string
	frequency int
}

func Top10(line string) []string {
	regTemplate := regexp.MustCompile(`([\.\,\"\;\:\!]+)|(-$)`)

	line = regTemplate.ReplaceAllString(line, "")
	splitLine := strings.Fields(line)
	if len(splitLine) == 1 {
		return splitLine
	}

	wordsBase := make(map[string]int, len(splitLine))
	for _, value := range splitLine {
		words := regTemplate.ReplaceAllString(value, "")
		if words != "" {
			wordsBase[strings.ToLower(words)]++
		}
	}

	wordsBaseRes := make([]Words, 0, len(wordsBase))
	for key, value := range wordsBase {
		wordsBaseRes = append(wordsBaseRes, Words{key, value})
	}

	sort.SliceStable(wordsBaseRes, func(i, j int) bool {
		if wordsBaseRes[i].frequency < wordsBaseRes[j].frequency {
			return false
		}
		if wordsBaseRes[i].frequency > wordsBaseRes[j].frequency {
			return true
		}
		return strings.Compare(wordsBaseRes[i].word, wordsBaseRes[j].word) == -1
	})

	resString := make([]string, 0, 10)
	for i := 0; i < len(wordsBaseRes) && i < 10; i++ {
		resString = append(resString, wordsBaseRes[i].word)
	}

	return resString
}
