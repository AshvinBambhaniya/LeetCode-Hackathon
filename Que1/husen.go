func longestCommonPrefix(strs []string) string {
	var answer string

	if len(strs) == 0 {
		return answer
	}
	if len(strs) <= 1 {
		return strs[0]
	}
	first := strs[0]

	for i, ch := range first {
		found := true
		for _, str := range strs[1:len(strs)] {

			if len(str) <= i {
				found = false
				break
			}
			if byte(ch) != str[i] {
				found = false
				break
			}

		}
		if !found {
			break
		} else {
			answer = fmt.Sprintf("%s%s", string(answer), string(ch))
		}
	}
	return answer
}