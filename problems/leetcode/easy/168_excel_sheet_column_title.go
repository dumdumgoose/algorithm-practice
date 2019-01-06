package easy

func convertToTitle(n int) string {
	if n == 0 {
		return ""
	}
	result := ""
	for {
		n -= 1
		result = string(rune(65+n%26)) + result
		if n < 26 {
			break
		}
		n /= 26
	}

	return result
}
