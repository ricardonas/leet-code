// Time complexity equals to O(n) linear.
func reverse(target string) string {
	runes := []rune(target)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
