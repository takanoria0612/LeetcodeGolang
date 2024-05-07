package GroupAnagrams

import "sort"

func GroupAnagrams(strs []string) [][]string {
	helper := make(map[string][]int)
	for idx, str := range strs {
		runes := []rune(str)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		sortedString := string(runes)
		helper[sortedString] = append(helper[sortedString], idx)
	}
	var res [][]string
	for _, idxs := range helper {
		words := make([]string, len(idxs))
		for i, idx := range idxs {
			words[i] = strs[idx]
		}
		res = append(res, words)
	}
	return res
}
