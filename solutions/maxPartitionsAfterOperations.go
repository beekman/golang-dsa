func maxPartitionsAfterOperations(s string, k int) int {
    n := len(s)
    memo := make(map[string]int)

    var dp func(idx, mask int, changed bool) int
    dp = func(idx, mask int, changed bool) int {
        if idx == n {
            return 0
        }

        key := fmt.Sprintf("%d_%d_%v", idx, mask, changed)
        if val, ok := memo[key]; ok {
            return val
        }

        distinctCount := countBits(mask)
        charBit := 1 << (s[idx] - 'a')
        result := 0

        // Option 1: Don't change current character
        if (mask & charBit) != 0 {
            // Character already in current partition
            result = dp(idx+1, mask, changed)
        } else if distinctCount < k {
            // Add character to current partition
            result = dp(idx+1, mask|charBit, changed)
        } else {
            // Start new partition
            result = 1 + dp(idx+1, charBit, changed)
        }

        // Option 2: Change current character (if not changed yet)
        if !changed {
            for newChar := 0; newChar < 26; newChar++ {
                if byte(newChar+'a') != s[idx] {
                    newBit := 1 << newChar
                    var option int

                    if (mask & newBit) != 0 {
                        option = dp(idx+1, mask, true)
                    } else if distinctCount < k {
                        option = dp(idx+1, mask|newBit, true)
                    } else {
                        option = 1 + dp(idx+1, newBit, true)
                    }

                    if option > result {
                        result = option
                    }
                }
            }
        }

        memo[key] = result
        return result
    }

    return dp(0, 0, false) + 1
}

func countBits(mask int) int {
    count := 0
    for mask > 0 {
        count += mask & 1
        mask >>= 1
    }
    return count
}
