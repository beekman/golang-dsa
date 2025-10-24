package main

import (
	"strconv"
)

/**
 * Helper function to check if a number is numerically balanced.
 * @param {int} n The number to check.
 * @return {bool} True if the number is balanced, false otherwise.
 */
func isNumericallyBalanced(n int) bool {
    // Convert the number to a string to iterate over its digits.
    s := strconv.Itoa(n)
    len := len(s)
    // Use a slice to store counts for digits 0-9.
    counts := make([]int, 10)

    // In Go, iterating a string gives runes (not bytes or chars).
    for _, r := range s {
        // Convert the rune '0'-'9' to an integer 0-9.
        digit := int(r - '0')

        // 1. Rule: Cannot contain 0.
        if digit == 0 {
            return false
        }

        // 2. Optimization: Same as in the JS solution.
        // If a digit `d` is > len, it's impossible.
        if digit > len {
            return false
        }

        counts[digit]++
    }

    // 3. Final check:
    // Go through each possible digit (1-9).
    for digit := 1; digit < 10; digit++ {
        // If this digit was present in the number (count > 0)...
        if counts[digit] > 0 {
            // ...its count must be exactly equal to its value.
            if counts[digit] != digit {
                return false
            }
        }
    }

    // If all checks passed, it's numerically balanced.
    return true
}

/**
 * @param {int} n
 * @return {int}
 */
func nextBeautifulNumber(n int) int {
    num := n + 1

    // Go's equivalent of `while(true)` is `for {}`.
    // We loop until we find and return the answer.
    for {
        if isNumericallyBalanced(num) {
            return num
        }
        num++
    }
}
