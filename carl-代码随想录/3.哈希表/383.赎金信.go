func canConstruct(ransomNote string, magazine string) bool {
	// 注意：rune是int32 的别名
	// string获得单个元素是rune
	myhash := make(map[rune]int)
	for _, c := range magazine {
		myhash[c]++
	}

	for _, c := range ransomNote {
		myhash[c]--
		if -1 == myhash[c] {
			return false
		}
	}
	return true
}
