package utils


// SplitStringByCharIndex splits the string by the given char at it's given occurrence. If the given occurrence is 0,
// both return values are the original string
func SplitStringByCharIndex(s string, c rune, index int) (string, string) {
	if index == 0 {
		return "", s
	}
	count := 0
	for i, char := range s {
		if char == c {
			count++
			if count == index {
				return  s[:i], s[i+1:]
			}
		}
	}
	return s, ""
}
