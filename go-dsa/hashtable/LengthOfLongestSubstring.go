package hashtable

func LengthOfLongestSubstring(s string) int {

	charMap := make(map[rune]int)
	left := 0
	maxLen := 0

	for right, char := range s {

		if pos, ok := charMap[char]; ok && pos >= left {
			left = pos + 1
		}

		charMap[char] = right

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}
