package main

// Renamed this function from 'makeArrayElementsEqualZero'
func countValidSelections(nums []int) int {
    n := len(nums)

    /**
     * Helper function to simulate the process for one (start, direction) pair.
     * startCurr: The starting index.
     * startDir: The starting direction (1 for right, -1 for left).
     * returns: True if the simulation results in all zeros, false otherwise.
     */
    check := func(startCurr int, startDir int) bool {
        // Create a copy of the array for this simulation
        numsCopy := make([]int, n)
        copy(numsCopy, nums)

        curr := startCurr
        dir := startDir

        // Loop until we go out of bounds
        for curr >= 0 && curr < n {
            if numsCopy[curr] > 0 {
                // Non-zero: Decrement, flip direction, and step
                numsCopy[curr]--
                dir *= -1 // dir = -dir
                curr += dir
            } else {
                // Zero: Fast-forward in the current direction
                for curr >= 0 && curr < n && numsCopy[curr] == 0 {
                    curr += dir
                }
            }
        }

        // After the simulation, check if all elements are zero
        for _, val := range numsCopy {
            if val != 0 {
                return false
            }
        }
        return true
    }

    validSelections := 0
    // Iterate through all possible starting positions
    for i := 0; i < n; i++ {
        if nums[i] == 0 {
            // Test starting at i, moving right
            if check(i, 1) {
                validSelections++
            }
            // Test starting at i, moving left
            if check(i, -1) {
                validSelections++
            }
        }
    }

    return validSelections
}
