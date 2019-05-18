package levenshtein

import (
	"fmt"
)

// LevenshteinDistance returns a metric used to determine the difference between two strings
// Following the algorithm: https://en.wikipedia.org/wiki/Levenshtein_distance
func Distance(s1, s2 string) int {
	m, n := len(s1), len(s2)
	// This algorithm starts from index 1, we need a dummy char at the front to
	// keep character access within bounds
	s1, s2 = fmt.Sprintf("_%s", s1), fmt.Sprintf("_%s", s2)
	matrix := make([][]int, m+1)
	for i := range matrix {
		matrix[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		matrix[i][0] = i
	}

	for j := 1; j <= n; j++ {
		matrix[0][j] = j
	}

	for j := 1; j <= n; j++ {
		for i := 1; i <= m; i++ {
			substitutionCost := 1
			if s1[i] == s2[j] {
				substitutionCost = 0
			}

			min := matrix[i-1][j] + 1
			if matrix[i][j-1]+1 < min {
				min = matrix[i][j-1] + 1
			}
			if matrix[i-1][j-1]+substitutionCost < min {
				min = matrix[i-1][j-1] + substitutionCost
			}
			matrix[i][j] = min
		}
	}
	return matrix[m][n]
}
