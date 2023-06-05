func divide(dividend int, divisor int) int {
    // Handle special cases
    if dividend == math.MinInt32 && divisor == -1 {
        return math.MaxInt32
    }

    // Determine the sign of the quotient
    isNeg := false
    if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
        isNeg = true
    }

    // Convert dividend and divisor to positive
    dividend = absValue(dividend)
    divisor = absValue(divisor)

    result := 0
    for dividend >= divisor {
        // Find the largest multiple of divisor that can be subtracted from dividend
        powerOfTwo := 1
        multiple := divisor
        for multiple*2 <= dividend {
            multiple *= 2
            powerOfTwo *= 2
        }

        // Subtract the multiple from dividend and add the corresponding quotient
        dividend -= multiple
        result += powerOfTwo
    }

    if isNeg {
        return -result
    }
    return result
}

func absValue(num int) int {
    if num < 0 {
        return -num
    }
    return num
}
