func isAnagram(s string, t string) bool {
	sMap := make(map[rune]int)
	tMap := make(map[rune]int)
	if len(s) != len(t) {
		return false
	}
	for _, sVlaue := range s {
		sMap[sVlaue]++
	}
	for _, tVlaue := range t {
		tMap[tVlaue]++
	}

	if len(sMap) != len(tMap) {
		return false
	}
	for k, v := range sMap {
		if tMap[k] != v {
			return false
		}
	}

	return true
}
