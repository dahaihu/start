package leetcode

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	mark := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		mark[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		mark[0][i] = i
	}
	for i := 1; i <= m; i++ {
		mark[i][0] = i
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				mark[i][j] = mark[i-1][j-1]
			} else {
				mark[i][j] = 1 + min(
					mark[i-1][j-1],
					min(mark[i][j-1], mark[i-1][j]),
				)
			}
		}
	}
	return mark[m][n]
}
