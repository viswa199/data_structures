package pattern

func Pattern(s string) []int {
	var lps []int
	lps = append(lps, 0)
	j := 0
	for i := 1; i < len(s); i++ {
		if string(s[j]) == string(s[i]) {
			lps = append(lps, j+1)
			j += 1
		} else {
			lps = append(lps, 0)
			j = 0
		}
	}
	return lps
}

func KMP(text, pattern string) []int {
	var kmp []int
	lps := Pattern(pattern)
	j := 0
	for i := 0; i < len(text); i++ {
	label:
		if string(text[i]) == string(pattern[j]) {
			j++
			if j == len(pattern) {
				kmp = append(kmp, i-j+1)
				j = lps[j-1]
			}
		} else {
			if j > 0 {
				j = lps[j-1]
				goto label
			}
		}
	}
	return kmp
}
