package main

// The refined version of the basic mayers diff algorithm is based on the concept of divide and conquer
// It operates in a way much similar to the quick sort algorithm by determining the middle point in each iteration
// Therefore, there is no DP matrix generated in this process

func getLinearMayersDiffOperations(src, dest []string) []operation {
	return linearMayerDiffHelper(src, dest, 0, 0)
}

func linearMayerDiffHelper(src, dest []string, offsetX, offsetY int) []operation {
	// Base case: one of the input string is empty
	// Both are empty --> no operation needed
	if len(src) == 0 && len(dest) == 0 {
		return []operation{}
	// src is empty --> all the remaining strings in dest are INSERT operations
	} else if len(src) == 0 {
		operations := make([]operation, len(dest))
		for i := 0; i < len(dest); i++ {
			operations[i] = operation{INSERT, offsetX, offsetY + i}
		}
		return operations
	// dest is empty --> all the remaining strings in src are DELETE operations
	} else if len(dest) == 0 {
		operations := make([]operation, len(src))
		for i := 0; i < len(src); i++ {
			operations[i] = operation{DELETE, offsetX + i, offsetY}
		}
		return operations
	}
	// None of the strings is empty --> find the middle snake
	headX, headY, tailX, tailY := findMiddleSnake(src, dest)
	// Divide the remaining strings into two parts
	operations := linearMayerDiffHelper(src[:headX], dest[:headY], offsetX, offsetY)
	operations = append(operations, linearMayerDiffHelper(src[tailX:], dest[tailY:], offsetX + tailX, offsetY + tailY)...)
	// Return the operations
	return operations
}

// Find the middle snake to divide the remaining strings into two parts
// Snake is the diagonal path following the INSERT or DELETE operations
// The length of snakes >= 0
func findMiddleSnake(src, dest []string) (headX, headY, tailX, tailY int) {
	// Use the same strategy as the basic mayers diff algorithm to find the middle snake
	// But search simultaneously from both ends of the strings
	m := len(src)
	n := len(dest)

	// The max possible operations needed to convert the strings
	max := (m + n + 1) / 2

	// Target K
	gapK := m - n

	// Inseated of using a DP matrix, use two slices to store the x values
	// Like what we have done in the getMayersDiffCount1D function
	forward := make([]int, 2*max+1)
	backward := make([]int, 2*max+1)

	// Initialize the search
	// The trigger the initial diagonal search
	backward[max + 1] = m

	// Scanner the strings forward
	for depth := 0; depth <= max; depth++ {
		// Fill in the new x values for branches
		for k := -depth; k <= depth; k += 2 {
			// Choose which branch to extend
			var x int
			var y int
			// If k == -depth, only k + 1 is available
			// If k != depth, then choose the branch with the greater x value for further outreach
			if k == -depth || (k != depth && forward[max + k - 1] < forward[max + k + 1]) {
				x = forward[max + k + 1]
				// SPECIAL CASE: After the operation, both string reach the end
				// PROBLEM: The division will fail --> infinite loop
				// SOLUTION: Use the previous snake as the middle snake
				y = x - k
				if x >= m && y >= n {
					return x, y-1, x, y-1
				}
			// If k == depth, only k - 1 is available
			} else {
				x = forward[max + k - 1] + 1
				// SPECIAL CASE HANDLER
				y = x - k
				if x >= m && y >= n {
					return x-1, y, x-1, y
				}
			}

			// Explore the diagonal: Try to find more matches
			// Moving along the diagonal will not cause any change in K and need no additional operations
			// Get the corresponding y value after the operation
			// k = x - y <==> y = x - k
			
			headX, headY := x, y
			// Follow the diagonal
			for x < m && y < n && trimRightSpace(src[x]) == trimRightSpace(dest[y]) {
				x++
				y++
			}

			// Save the x value after the operation and diagonal search
			// The diagonal searched blocks is called a *snake* (length >= 0)
			forward[max + k] = x

			// Check if any overlapping is found
			// Since len(src) - len(dest) = gapK, the overlapping will be found when the gapK is odd
			// gapK indicates the parity of the number of operations needed to convert the strings
			if gapK % 2 == 1 && gapK - depth < k && k < gapK + depth {
				// If the overlapping is found, return the snake
				if forward[max + k] >= backward[max + gapK - k] {
					return headX, headY, x, y
				}
			}
		}

		// Fill in the new x values for branches
		for k := -depth; k <= depth; k += 2 {
			// Choose which branch to extend
			var x int
			// If k == -depth, only k + 1 is available
			// If k != depth, then choose the branch with the greater x value for further outreach
			if k == -depth || (k != depth && backward[max + k - 1] > backward[max + k + 1]) {
				x = backward[max + k + 1]
			// If k == depth, only k - 1 is available
			} else {
				// Moving left
				x = backward[max + k - 1] - 1
			}

			// Explore the diagonal: Try to find more matches
			// Moving along the diagonal will not cause any change in K and need no additional operations
			// Get the corresponding y value after the operation
			// The original equation for forward extension is not available here
			y := x + k - gapK
			tailX, tailY := x, y
			// Follow the diagonal reversely
			for x > 0 && y > 0 && trimRightSpace(src[x-1]) == trimRightSpace(dest[y-1]) {
				x--
				y--
			}

			// Save the x value after the operation and diagonal search
			// The diagonal searched blocks is called a *snake* (length >= 0)
			backward[max + k] = x

			// Check if any overlapping is found
			// Since len(src) - len(dest) = gapK, the overlapping will be found when the gapK is even
			// gapK indicates the parity of the number of operations needed to convert the strings
			if gapK % 2 == 0 && k >= gapK - depth && k <= gapK + depth {
				// If the overlapping is found, return the snake
				if forward[max + gapK - k] >= backward[max + k] {
					return x, y, tailX, tailY
				}
			}
		}
	}

	// There must be a middle snake in the strings
	// The default return will never be used
	return -1, -1, -1, -1
}