func firstUniqChar(s string) byte {
	myhash := make(map[byte]int)
	for _, c := range s {
		myhash[byte(c)]++
	}
	for _, c := range s {
		if 1 == myhash[byte(c)] {
			return byte(c)
		}
	}
	return ' '
}
