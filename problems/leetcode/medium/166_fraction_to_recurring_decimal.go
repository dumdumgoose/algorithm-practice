package medium

import (
	"math"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	result := ""
	if (numerator > 0) != (denominator > 0) {
		result += "-"
	}

	numVal := int(math.Abs(float64(numerator)))
	denomVal := int(math.Abs(float64(denominator)))
	integral := numVal / denomVal
	remainder := numVal % denomVal
	result += strconv.Itoa(integral)
	if remainder == 0 {
		return result
	}
	result += "."
	remainder *= 10
	fractionMap := make(map[int]int)
	index := len(result)
	for remainder != 0 {
		if _, ok := fractionMap[remainder]; ok {
			insert := fractionMap[remainder]
			result = result[:insert] + "(" + result[insert:] + ")"
			break
		}
		result += strconv.Itoa(remainder / denomVal)
		fractionMap[remainder] = index
		remainder = remainder % denomVal * 10
		index++
	}
	return result
}
