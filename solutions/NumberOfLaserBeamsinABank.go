// This 'package' declaration is required
package main

// We only import packages we actually use
import (
	"strings"
)

/**
 * This is the function the platform will find and test.
 */
func numberOfBeams(bank []string) int {
	var totalBeams int = 0
	var prevDeviceCount int = 0

	// Loop through each row in the bank slice
	for _, row := range bank {
		// Count the devices ('1's) in the current row
		currentDeviceCount := strings.Count(row, "1")

		// If this row has devices...
		if currentDeviceCount > 0 {
			// Add beams connecting to the previous device-row
			totalBeams += prevDeviceCount * currentDeviceCount

			// This row now becomes the 'previous' row for the next iteration
			prevDeviceCount = currentDeviceCount
		}
		// If currentDeviceCount is 0, we do nothing and
		// 'prevDeviceCount' keeps its old value, carrying
		// over to the next row that has devices.
	}

	// Return the final count
	return totalBeams
}
