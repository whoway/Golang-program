func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// 注意：rune是int32 的别名
	// string获得单个元素是rune
	myhash := make(map[rune]int)
	for _, c := range s {
		myhash[c]++
	}

	for _, c := range t {
		myhash[c]--
		if -1 == myhash[c] {
			return false
		}
	}
	return true
}
