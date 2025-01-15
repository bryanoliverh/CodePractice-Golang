func minimumLength(s string) int {
	count := make(map[rune]int)
	for _, char := range s {
		count[char]++
	}
	minus := 0
	for _, freq := range count {
		for freq >= 3 {
			minus += 2
			freq -= 2
		}
	}
	return len(s) - minus
}