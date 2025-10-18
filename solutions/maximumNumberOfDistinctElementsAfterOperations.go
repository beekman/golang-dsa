package main

import (
	"math"
	"sort"
)

// max returns the maximum of two int64 values.
func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// Function renamed to maxDistinctElements to match the expected LeetCode signature.
func maxDistinctElements(nums []int, k int) int {
	// Convert to int64 for calculations to prevent potential overflow,
	// as nums[i] and k can be up to 10^9.
	n := len(nums)
	nums64 := make([]int64, n)
	for i := 0; i < n; i++ {
		nums64[i] = int64(nums[i])
	}

	// 1. Sort the array.
	sort.Slice(nums64, func(i, j int) bool {
		return nums64[i] < nums64[j]
	})

	k64 := int64(k)
	distinctCount := 0

	// maxUsed tracks the largest final value chosen so far.
	maxUsed := int64(math.MinInt64)

	// 2. Greedy assignment.
	for _, num := range nums64 {
		// Calculate the range of possible final values [L, R].
		L := num - k64 // Smallest possible new value
		R := num + k64 // Largest possible new value

		// The new value must be strictly greater than maxUsed.
		// The smallest possible value we can use is maxUsed + 1.
		candidate := maxUsed + 1

		// The chosen new value must satisfy: newVal >= L AND newVal > maxUsed.
		// The smallest value satisfying both is max(L, candidate).
		newVal := max(L, candidate)

		// Check if this newVal is within the element's allowed range [L, R].
		if newVal <= R {
			// A distinct value is found and used.
			maxUsed = newVal
			distinctCount++
		}
	}

	return distinctCount
}
