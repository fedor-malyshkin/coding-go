package p438

/*
https://leetcode.com/problems/find-all-anagrams-in-a-string/
*/
func findAnagrams(s string, p string) []int {
	pMap := map[uint8]int{}
	sMap := map[uint8]int{}
	for i := 0; i < len(p); i++ {
		pMap[p[i]] = pMap[p[i]] + 1
	}
	if len(p) > len(s) {
		return []int{}
	}
	var (
		st = 0
		en = len(p) - 1
	)
	for i := st; i <= en; i++ {
		sMap[s[i]] = sMap[s[i]] + 1
	}
	res := []int{}
	for en < len(s) {
		if match(pMap, sMap) {
			res = append(res, st)
		}
		sMap[s[st]] = sMap[s[st]] - 1
		st++
		en++
		if en < len(s) {
			sMap[s[en]] = sMap[s[en]] + 1
		}
	}
	return res

}

func match(pMap, sMap map[uint8]int) bool {
	for k, v := range pMap {
		sv := sMap[k]
		if v != sv {
			return false
		}
	}
	for k, v := range sMap {
		pv := pMap[k]
		if v != pv {
			return false
		}
	}
	return true
}
