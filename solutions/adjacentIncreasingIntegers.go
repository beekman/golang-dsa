package main

// The unused 'math' import is removed.
// We rename the function to 'maxIncreasingSubarrays' to resolve the 'undefined' error
// from the test runner.

/**
 * The original function logic:
 * func maxAdjacentIncreasingSubarrays(nums []int) int
 * is renamed to:
 * func maxIncreasingSubarrays(nums []int) int
 */
func maxIncreasingSubarrays(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	// 1. Precompute LIS_end: Longest strictly increasing subarray ending at index i
	LIS_end := make([]int, n)
	for i := range LIS_end {
		LIS_end[i] = 1
	}

	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			LIS_end[i] = LIS_end[i-1] + 1
		}
	}

	// canFind checks if two adjacent strictly increasing subarrays of length k exist.
	canFind := func(k int) bool {
		if k*2 > n {
			return false
		}
		// a is the starting index of the first subarray. a <= n - 2k.
		for a := 0; a <= n-2*k; a++ {
			end1 := a + k - 1   // End index of the first subarray
			end2 := a + 2*k - 1 // End index of the second subarray

			// Check if the first subarray nums[a..end1] is strictly increasing (length k).
			isFirstIncreasing := LIS_end[end1] >= k

			// Check if the second subarray nums[a+k..end2] is strictly increasing (length k).
			isSecondIncreasing := LIS_end[end2] >= k

			if isFirstIncreasing && isSecondIncreasing {
				return true
			}
		}
		return false
	}

	// Binary Search for the maximum k.
	low := 1
	high := n / 2 // Maximum possible length for k
	maxK := 0

	for low <= high {
		mid := low + (high-low)/2
		if canFind(mid) {
			// mid is possible, try for a larger k
			maxK = mid
			low = mid + 1
		} else {
			// mid is too large, search for a smaller k
			high = mid - 1
		}
	}

	return maxK
}
