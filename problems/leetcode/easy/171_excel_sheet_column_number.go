package easy

func titleToNumber(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		result = 26*result + int(s[i]) - 64
	}
	return result
}
