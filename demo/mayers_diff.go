package main

// The DP matrix in the mayers diff algorithm is different from the one in the basic diff algorithm
// It is a matrix using Depth (D) as the x-axis and the change of length (K) as the y-axis
func getMayersDiffDP(src, dest []string) *[][]int {
	m := len(src)
	n := len(dest)

	// The max possible operations needed to convert the strings
	// Worst scenario: all the elements in src are deleted, then all the elements in dest are inserted
	max := m + n

	// Current: the current row of processing
	current := make([]int, 2*max+1)

	// History: the DP matrix
	history := [][]int{}

	// Perform the breadth-first search
	// Discover two new branches at each step
	for depth := 0; depth <= max; depth++ {
		// Fill in the new x values for branches
		for k := -depth; k <= depth; k += 2 {
			// Choose which branch to extend
			var x int
			// If k == -depth, only k + 1 is available
			// If k != depth, then choose the branch with the greater x value for further outreach
			if k == -depth || (k != depth && current[max + k - 1] < current[max + k + 1]) {
				x = current[max + k + 1]
			// If k == depth, only k - 1 is available
			} else {
				x = current[max + k - 1] + 1
			}

			// Explore the diagonal: Try to find more matches
			// Moving along the diagonal will not cause any change in K and need no additional operations
			// Get the corresponding y value after the operation
			// k = x - y <==> y = x - k
			y := x - k
			// Follow the diagonal
			for x < m && y < n && trimRightSpace(src[x]) == trimRightSpace(dest[y]) {
				x++
				y++
			}

			// Save the x value after the operation and diagonal search
			// The combination of the operation and diagonal search is called a *snake*
			current[max + k] = x

			// If the search is done, save the history and return
			if x >= m && y >= n {
				temp := make([]int, 2*depth+1)
				copy(temp, current[max-depth:max+depth+1])
				history = append(history, temp)
				return &history
			}
		}
		// Save the result of this depth
		temp := make([]int, 2*depth+1)
		copy(temp, current[max-depth:max+depth+1])
		history = append(history, temp)
	}

	return &history
}

func getMayersDiffOperations(src, dest []string) opStack {
	// Get the search history
	history := *getMayersDiffDP(src, dest)

	// Find the final K value
	// what is the K value of the lower-right corner of the edit graph
	// k = x - y ==> k = len(src) - len(dest)
	k := len(src) - len(dest)

	// Start from the last searched depth
	// Perform the reverse process of the breadth-first search
	operations := opStack{}
	for depth := len(history) - 1; depth > 0; depth-- {
		// Determine the parent branch
		// In other words, determine the K value of the parent depth
		previous := history[depth-1]
		var insert bool
		if k == -depth || (k != depth && previous[depth - 1 + k - 1] < previous[depth - 1 + k + 1]) {
			k = k + 1
			insert = true
		} else {
			k = k - 1
			insert = false
		}

		// Get the x and y values of the operation
		// k = x - y ==> y = x - k
		x := previous[depth - 1 + k]
		y := x - k

		// Record the operation
		if insert {
			operations.push(operation{INSERT, x, y})
		} else {
			operations.push(operation{DELETE, x, y})
		}
	}

	return operations
}

/** The following is not used in getting diff */
// Just serve as a reference for the ways to reduce space complexity

func getMayersDiffCount(src, dest []string) int {
	// Get the dp array
	dp := *getMayersDiffDP(src, dest)

	return len(dp) - 1
}

func getMayersDiffCount1D(src, dest []string) int {
	m := len(src)
	n := len(dest)

	// The max possible operations needed to convert the strings
	// Worst scenario: all the elements in src are deleted, then all the elements in dest are inserted
	max := m + n

	// Current: the current row of processing
	current := make([]int, 2*max+1)

	// No need of history since the path is not needed

	// Perform the breadth-first search
	// Discover two new branches at each step
	for depth := 0; depth <= max; depth++ {
		// Fill in the new x values for branches
		for k := -depth; k <= depth; k += 2 {
			// Choose which branch to extend
			var x int
			// If k == -depth, only k + 1 is available
			// If k != depth, then choose the branch with the greater x value for further outreach
			if k == -depth || (k != depth && current[max + k - 1] < current[max + k + 1]) {
				x = current[max + k + 1]
			// If k == depth, only k - 1 is available
			} else {
				x = current[max + k - 1] + 1
			}

			// Explore the diagonal: Try to find more matches
			// Moving along the diagonal will not cause any change in K and need no additional operations
			// Get the corresponding y value after the operation
			// k = x - y <==> y = x - k
			y := x - k
			// Follow the diagonal
			for x < m && y < n && trimRightSpace(src[x]) == trimRightSpace(dest[y]) {
				x++
				y++
			}

			// Save the x value after the operation and diagonal search
			// The diagonal searched blocks is called a *snake* (length >= 0)
			current[max + k] = x

			// If the search is done, return the depth
			if x >= m && y >= n {
				return depth
			}
		}
	}

	return max
}