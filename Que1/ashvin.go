func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for {
			if len(prefix) > len(strs[i]) {
				prefix = prefix[:len(prefix)-1]
				if prefix == "" {
					return ""
				}
			} else {
				break
			}
		}
		for {
			if strs[i][:len(prefix)] == prefix {
				break
			} else {
				prefix = prefix[:len(prefix)-1]
				if prefix == "" {
					retSurn ""
				}
			}
		}
	}
	return prefix
}