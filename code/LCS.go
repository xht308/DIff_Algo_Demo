package main

/** The original solution to the LCS problem is not used in Diff */
// The algorithms are kept here for reference

// Get the dynamic programming array for the longest common subsequence
func getLCSDP(text1, text2 string) *[][]int {
	m := len(text1)
	n := len(text2)

	// Dynamic Programming: Create a 2D array to store the lengths of common subsequences
	// dp[i][j] = length of longest common subsequence of text1[0..i-1] and text2[0..j-1]
	// dp[m][n] = length of longest common subsequence of text1[0..m-1] and text2[0..n-1]
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	// Fill the dp array
	// If the characters match, then append the character to the common subsequence
	// If the characters don't match, then take the max of the two subsequences
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// After all the iterations, the last element of the dp array will contain the length of the longest common subsequence
	return &dp
}

// Get the longest common subsequence
func getLCS(text1, text2 string) string {
	// Get the dp array
	dp := *getLCSDP(text1, text2)

	m := len(text1)
	n := len(text2)

	// Start from the bottom right corner of the dp array
	// If the characters match, then append the character to the common subsequence
	// If the characters don't match, then go in the direction of the greater subsequence
	// Notably, appending the character to the common subsequence cause a diagonal movement
	lcs := ""
	for m > 0 && n > 0 {
		if text1[m-1] == text2[n-1] {
			lcs = string(text1[m-1]) + lcs
			m--
			n--
		} else if dp[m-1][n] > dp[m][n-1] {
			m--
		} else {
			n--
		}
	}

	return lcs
}

/** The following is not used in getting diff */
// Just serve as a reference for the ways to reduce space complexity

func getLCSLengthRecursive(text1, text2 string) int {
	m := len(text1)
	n := len(text2)

	// Base case: If either of the strings is empty, then the length of the longest common subsequence is 0
	if m == 0 || n == 0 {
		return 0
	}

	// Recursive case: If the characters match, then the length of the longest common subsequence is 1 + the length of the longest common subsequence of the rest of the strings
	// If the characters don't match, then the length of the longest common subsequence is the max of the two possible subsequences
	if text1[0] == text2[0] {
		return getLCSLengthRecursive(text1[1:], text2[1:]) + 1
	} else {
		return max(getLCSLengthRecursive(text1[1:], text2), getLCSLengthRecursive(text1, text2[1:]))
	}
}

func getLCSLength(text1, text2 string) int {
	// Get the dp array
	dp := *getLCSDP(text1, text2)

	// After all the iterations, the last element of the dp array will contain the length of the longest common subsequence
	return dp[len(text1)][len(text2)]
}

func getLCSLength1D(text1, text2 string) int {
	m := len(text1)
	n := len(text2)

	// Dynamic Programming: Create a 1D array to store the lengths of common subsequences
	// Use a 1D array (rolling array) instead of a 2D array to save space
	dp := make([]int, n+1)

	// Fill the dp array
	// If the characters match, then append the character to the common subsequence
	// If the characters don't match, then take the max of the two subsequences
	for i := 1; i <= m; i++ {
		prev := 0
		for j := 1; j <= n; j++ {
			temp := dp[j]
			if text1[i-1] == text2[j-1] {
				dp[j] = prev + 1
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			prev = temp
		}
	}

	// After all the iterations, the last element of the dp array will contain the length of the longest common subsequence
	return dp[n]
}