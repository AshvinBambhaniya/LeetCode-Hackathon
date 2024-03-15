func longestCommonPrefix(strs []string) string {
	var cP string
	var currentChar string
	var j int
	for _, i := range strs[0] {
		currentChar = string(i)
		var l int

		for _, k := range strs {
			if j == len(k) {
				return cP
			}
			if string(k[j]) == currentChar {
				l = l + 1
			} else {
				break
			}
		}
		// fmt.Print(l)
		if l == len(strs) {
			cP = cP + currentChar
		} else {
			return cP
		}
		j = j + 1
	}
	return cP
}

// solved at 6:50.26 pm