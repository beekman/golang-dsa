package main

import (
	"sort"
)

// Go 1.20 and earlier need explicit min/max
func min(a, b int) int {
if a < b {
	return a
}
return b
}

func max(a, b int) int {
if a > b {
	return a
}
return b
}

func maxFrequency(nums []int, k_int int, numOperations int) int {
n := len(nums)
// Use int64 for k and all calculations to prevent overflow
// (e.g., target+k can be 10^9 + 10^9 + 10^9 = 3e9, which overflows int32)
k := int64(k_int)

// Convert nums to []int64 for safe calculations
nums64 := make([]int64, n)
for i, num := range nums {
	nums64[i] = int64(num)
}

// Step 1: Sort the array. O(N log N)
sort.Slice(nums64, func(i, j int) bool {
	return nums64[i] < nums64[j]
})

// Step 2: Create a set of candidate targets
// We use a map as a set in Go
candidates := make(map[int64]bool)
for _, num := range nums64 {
	candidates[num] = true
	candidates[num-k] = true
	candidates[num+k] = true
}

maxFreq := 0

// Step 3: Check each candidate. O(N * log N)
for target := range candidates {
	// Use sort.Search, which is Go's lower_bound

	// Find g(T): count of elements == target
	// lower_bound(target)
	l_g := sort.Search(n, func(i int) bool {
		return nums64[i] >= target
	})
	// upper_bound(target)
	r_g := sort.Search(n, func(i int) bool {
		return nums64[i] > target
	})
	alreadyEqual := r_g - l_g

	// Find f(T): count of elements in [target - k, target + k]
	// lower_bound(target - k)
	l_f := sort.Search(n, func(i int) bool {
		return nums64[i] >= (target - k)
	})
	// upper_bound(target + k)
	r_f := sort.Search(n, func(i int) bool {
		return nums64[i] > (target + k)
	})
	totalInRange := r_f - l_f

	// Elements that can become target but aren't already
	canBecome := totalInRange - alreadyEqual

	// (those already equal) + (those we can change, limited by ops)
	currentFreq := alreadyEqual + min(canBecome, numOperations)
	maxFreq = max(maxFreq, currentFreq)
}

return maxFreq
}
