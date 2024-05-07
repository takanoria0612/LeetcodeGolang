package GroupAnagrams

type CharTable [26]byte

func (t *CharTable) Push(value byte) {
	t[value-'a']++
}

func GroupAnagrams3(strs []string) [][]string {
	group := make(map[CharTable][]string)
	for i := range strs {
		var key CharTable
		for j := range strs[i] {
			key.Push(strs[i][j])
		}
		group[key] = append(group[key], strs[i])
	}
	return getMapValues(group)
}

func getMapValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}
