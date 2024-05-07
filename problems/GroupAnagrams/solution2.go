package GroupAnagrams

import (
	"sort"
	"strings"
)

func GroupAnagrams2(strs []string) [][]string {
	anagramMap := make(map[string][]string)
	var result [][]string

	for _, s := range strs {
		sortedS := sortString(s)
		anagramMap[sortedS] = append(anagramMap[sortedS], s)
	}

	for _, value := range anagramMap {
		result = append(result, value)
	}

	return result
}

func sortString(s string) string {
	sSlice := strings.Split(s, "")
	sort.Strings(sSlice)
	return strings.Join(sSlice, "")
}
