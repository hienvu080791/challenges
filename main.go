package main

func main() {
	//fmt.Println(grayCode(2))
	//fmt.Println(sumOfDistancesInTree(2, [][]int{{0, 1}}))
	//fmt.Println(findLength([]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}))
}

func findLength(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	maxLength := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxLength {
					maxLength = dp[i][j]
				}
			}
		}
	}
	return maxLength
}

func sumOfDistancesInTree(n int, edges [][]int) []int {
	switch {
	case n < 1 || n > 30000:
		return nil
	case n == 1:
		if len(edges) > 0 {
			return nil
		}
		return []int{0}
	case len(edges) != n-1:
		return nil
	}

	graph := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	count := make([]int, n)
	ans := make([]int, n)

	var dfs1 func(node, parent int)
	dfs1 = func(node, parent int) {
		count[node] = 1
		for _, child := range graph[node] {
			if child == parent {
				continue
			}
			dfs1(child, node)
			count[node] += count[child]
			ans[node] += ans[child] + count[child]
		}
	}
	var dfs2 func(node, parent int)
	dfs2 = func(node, parent int) {
		for _, child := range graph[node] {
			if child == parent {
				continue
			}
			ans[child] = ans[node] - count[child] + (n - count[child])
			dfs2(child, node)
		}
	}
	dfs1(0, -1)
	dfs2(0, -1)
	return ans
}

func grayCode(n int) []int {
	if n < 1 || n > 16 {
		return nil
	}
	count := 1 << n
	result := make([]int, count)
	for i := 0; i < count; i++ {
		result[i] = i ^ (i >> 1)
	}
	return result
}
