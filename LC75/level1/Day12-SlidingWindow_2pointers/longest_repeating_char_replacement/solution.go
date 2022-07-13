package main

func characterReplacements(s string, k int) int {
	count := make(map[byte]int, 26)

	res := 0

	left, maxFreq := 0, 0

	for right := 0; right < len(s); right++ {
		count[s[right]]++

		maxFreq = max(maxFreq, count[s[right]])

		if (right-left+1)-maxFreq > k {
			count[s[left]]--
			left++
		}

		res = max(res, right-left+1)
	}

	return res
}

func max(a, b int) int {

	if a > b {
		return a
	}

	return b
}
