func isAnagram(s string, t string) bool {
	sArr := strings.Split(s, "")
	tArr := strings.Split(t, "")
	sort.Strings(sArr)
	sort.Strings(tArr)
	//比较字符数组是否相同的方法应该使用reflect.DeepEqual函数
	return reflect.DeepEqual(sArr, tArr)
}
