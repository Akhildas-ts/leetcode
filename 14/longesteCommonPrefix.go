package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i, char := range strs[0] {

		for _, str := range strs[1:] {
			if i >= len(str) || rune(str[i]) != char {

				return strs[0][:i]
			}
		}
	}

	return strs[0]
}
