package main

func getBasicDiffDP(src, dest []string) *[][]int {
	m := len(src)
	n := len(dest)

	// Dynamic Programming: Create a 2D array to store the number of operations needed to convert the strings
	// dp[i][j] = number of operations needed to convert text1[0..i-1] to text2[0..j-1]
	// Initialize the first row and column to 0, 1, 2, 3, ...
	dp := make([][]int, m+1)
	for x := 0; x <= m; x++ {
		temp := make([]int, n+1)
		temp[0] = x
		dp[x] = temp
	}
	for y := 0; y <= n; y++ {
		dp[0][y] = y
	}

	// Fill the dp array
	// If the characters match, no operations needed to keep the sequence the same
	// If the characters don't match, then take an additional operation to change the character
	for x := 1; x <= m; x++ {
		for y := 1; y <= n; y++ {
			if trimRightSpace(src[x-1]) == trimRightSpace(dest[y-1]) {
				dp[x][y] = dp[x-1][y-1]
			} else {
				dp[x][y] = min(dp[x-1][y], dp[x][y-1]) + 1
			}
		}
	}

	return &dp
}

func getBasicDiffOperations(src, dest []string) opStack {
	// Get the dp array
	dp := *getBasicDiffDP(src, dest)

	x := len(src)
	y := len(dest)
	
	// Start from the bottom right corner of the dp array
	// If the characters match, then append the character to the common subsequence
	// If the characters don't match, then go in the direction of the greater subsequence
	// Notably, appending the character to the common subsequence cause a diagonal movement
	operations := opStack{}
	for x > 0 && y > 0 {
		if trimRightSpace(src[x-1]) == trimRightSpace(dest[y-1]) {
			x--
			y--
		} else if dp[x-1][y] < dp[x][y-1] {
			operations.push(operation{DELETE, x - 1, y})
			x--
		} else {
			operations.push(operation{INSERT, x, y - 1})
			y--
		}
	}

	// Traverse the remaining characters
	// All are insertions and deletions done at the beginning of the string
	for x > 0 {
		operations.push(operation{DELETE, x - 1, 0})
		x--
	}
	for y > 0 {
		operations.push(operation{INSERT, 0, y - 1})
		y--
	}

	return operations
}

/** The following is not used in getting diff */
// Just serve as a reference for the ways to reduce space complexity

func getBasicDiffCount(src, dest []string) int {
	// Get the dp array
	dp := *getBasicDiffDP(src, dest)

	return dp[len(src)][len(dest)]
}

func getBasicDiffCount1D(src, dest []string) int {
	m := len(src)
	n := len(dest)

	// Dynamic Programming: Create a 1D array to store the number of operations needed to convert the strings
	// Use a 1D array (rolling array) instead of a 2D array to save space
	dp := make([]int, n+1)
	for y := 0; y <= n; y++ {
		dp[y] = y
	}

	// Fill the dp array
	// If the characters match, no operations needed to keep the sequence the same
	// If the characters don't match, then take an additional operation to change the character
	for x := 1; x <= m; x++ {
		prev := dp[0]
		dp[0] = x
		for y := 1; y <= n; y++ {
			temp := dp[y]
			if trimRightSpace(src[x-1]) == trimRightSpace(dest[y-1]) {
				dp[y] = prev
			} else {
				dp[y] = min(dp[y], dp[y-1]) + 1
			}
			prev = temp
		}
	}

	return dp[n]
}