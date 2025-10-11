package main

import (
	"sort"
)

// max returns the larger of two integers.
func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

/**
 * @param {[]int} power
 * @return {int}
 */
func maximumTotalDamage(power []int) int64 {
    // 1. Preprocessing: Calculate total damage for each unique power value.
    // Use a map to store total damage, since power values can be up to 10^9.
    // We use int64 for damage sums to prevent potential overflow, as the sum
    // can exceed 2^32 (10^5 elements * 10^9 max value).
    damageMap := make(map[int]int64)
    for _, p := range power {
        damageMap[p] += int64(p)
    }

    // 2. Get unique powers and sort them.
    uniquePowers := make([]int, 0, len(damageMap))
    for p := range damageMap {
        uniquePowers = append(uniquePowers, p)
    }
    sort.Ints(uniquePowers)

    n := len(uniquePowers)
    if n == 0 {
        return 0
    }

    // 3. Dynamic Programming
    // dp[i] stores the maximum total damage considering spells with power up to uniquePowers[i].
    // Since uniquePowers has at most 10^5 elements, an array is efficient here.
    dp := make([]int64, n)

    // dp[0] base case: The max damage is simply the damage from the smallest power.
    dp[0] = damageMap[uniquePowers[0]]

    for i := 1; i < n; i++ {
        p := uniquePowers[i]
        currentDamage := damageMap[p]

        // --- Choice 1: Cast spells with damage p ---
        // Damage gained is currentDamage + max damage from an allowed previous state.
        // The previous state must correspond to a power <= p - 3.

        var maxPreviousDamage int64 = 0

        // Search backwards for the largest index j < i such that uniquePowers[j] <= p - 3.
        // Since the array is sorted, the first one we find is the largest.
        j := i - 1
        for j >= 0 {
            if uniquePowers[j] <= p - 3 {
                // dp[j] holds the max damage up to uniquePowers[j].
                maxPreviousDamage = dp[j]
                break
            }
            j--
        }

        damageWithP := currentDamage + maxPreviousDamage

        // --- Choice 2: Skip spells with damage p ---
        // Max damage is simply the max damage up to uniquePowers[i-1].
        damageWithoutP := dp[i-1]

        // DP transition: Choose the maximum of the two options.
        dp[i] = max(damageWithP, damageWithoutP)
    }

    // The result is the max damage considering all unique powers.
    return dp[n-1]
}
