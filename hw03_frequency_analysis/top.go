package hw03frequencyanalysis

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func Top10(line string) []string {

	line = regexp.MustCompile(`[\.\,\"\;\:\!]`).ReplaceAllString(line, "")
	splitLine := strings.Fields(line)

	wordsBase := make(map[string]int, len(splitLine))

	for i := 0; i < len(splitLine); i++ {
		if value, ok := wordsBase[splitLine[i]]; ok {
			wordsBase[splitLine[i]] = value + 1
		} else {
			wordsBase[splitLine[i]] = 1
		}
	}

	keys := make([]string, 0, len(wordsBase))
	for k := range wordsBase {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return wordsBase[keys[i]] > wordsBase[keys[j]]
	})

	resString := make([]string, 0)
	count := 10
	for _, k := range keys {
		if count != 0 {
			resString = append(resString, k)
			fmt.Println(k, wordsBase[k])
			count--
		} else {
			break
		}
	}

	return resString
}
