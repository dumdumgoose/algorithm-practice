package medium

import "math"

/*
https://leetcode.com/problems/prime-palindrome/
Find the smallest prime palindrome greater than or equal to N.

Recall that a number is prime if it's only divisors are 1 and itself, and it is greater than 1.

For example, 2,3,5,7,11 and 13 are primes.

Recall that a number is a palindrome if it reads the same from left to right as it does from right to left.

For example, 12321 is a palindrome.

Example 1:
Input: 6
Output: 7

Example 2:
Input: 8
Output: 11

Example 3:
Input: 13
Output: 101


Note:
1 <= N <= 10^8
The answer is guaranteed to exist and be less than 2 * 10^8.
*/

func primePalindrome(n int) int {
	for {
		if n == reverseInt(n) && isPrime(n) {
			return n
		}
		n++
		// there's prime palindrome in this range, just skip, found this trick in solution lol
		if n > 10000000 && n < 100000000 {
			n = 100000000
		}
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	r := int(math.Sqrt(float64(n)))
	for d := 2; d <= r; d++ {
		if n%d == 0 {
			return false
		}
	}

	return true
}

func reverseInt(n int) int {
	r := 0

	for n > 0 {
		r = r*10 + n%10
		n /= 10
	}

	return r
}
