package pkg

// Function to calculate the Levenshtein distance between two words
func levenshteinDistance(s1, s2 string) int {
	r1, r2 := []rune(s1), []rune(s2)
	len1, len2 := len(r1), len(r2)

	// Create a matrix to store partial distances
	matrix := make([][]int, len1+1)
	for i := range matrix {
		matrix[i] = make([]int, len2+1)
	}

	// Initialize the matrix
	for i := 0; i <= len1; i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= len2; j++ {
		matrix[0][j] = j
	}

	// Fill the matrix with minimum distances
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			cost := 0
			if r1[i-1] != r2[j-1] {
				cost = 1
			}
			matrix[i][j] = min(min(matrix[i-1][j]+1, matrix[i][j-1]+1), matrix[i-1][j-1]+cost)
		}
	}

	// The last cell of the matrix contains the Levenshtein distance
	return matrix[len1][len2]
}

// Helper function to calculate the minimum of three integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Find function to find words in the dictionary with a maximum distance from the target word
func Find(dictionary []string, maxDistance int, targetWord string) []string {
	var result []string

	for _, word := range dictionary {
		distance := levenshteinDistance(targetWord, word)
		if distance <= maxDistance {
			result = append(result, word)
		}
	}

	return result
}
