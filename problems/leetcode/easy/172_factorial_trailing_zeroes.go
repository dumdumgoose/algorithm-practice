package easy

import "math"

/*
This is my original solution, it doesn't has any logic issue, just looking for how many 2 and 5 pairs in the numbers,
but since leetcode has very strict time limit so this one doesn't pass the test for 1808548329.
*/
func trailingZeroes(n int) int {
	if n < 5 {
		return 0
	}

	counter5 := 0
	counter2 := 0

	for i := 2; i <= n; i++ {
		for j := i; j > 1; {
			if j%5 == 0 {
				counter5++
				j /= 5
			} else if j%2 == 0 {
				counter2++
				j /= 2
			} else {
				break
			}
		}
	}

	return int(math.Min(float64(counter5), float64(counter2)))
}

func trailingZeroesLoop(n int) int {
	result := 0
	for n > 0 {
		result += n / 5
		n /= 5
	}
	return result
}

func trailingZeroesRecursive(n int) int {
	if n < 5 {
		return 0
	}

	if n < 10 {
		return 1
	}

	return n/5 + trailingZeroesRecursive(n/5)
}
