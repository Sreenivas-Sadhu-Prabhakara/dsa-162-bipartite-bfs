package main

import "fmt"

// IsBipartite reports whether the undirected graph can be 2-colored so no
// edge joins same-colored vertices — BFS alternates colors level by level
// and fails on any edge between equals.
func IsBipartite(n int, edges [][2]int) bool {
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	color := make([]int, n) // 0 = uncolored, 1 and 2 the two colors
	for s := 0; s < n; s++ {
		if color[s] != 0 {
			continue
		}
		color[s] = 1
		queue := []int{s}
		for len(queue) > 0 {
			v := queue[0]
			queue = queue[1:]
			for _, w := range adj[v] {
				if color[w] == 0 {
					color[w] = 3 - color[v] // the other color
					queue = append(queue, w)
				} else if color[w] == color[v] {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	fmt.Println(IsBipartite(4, [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}}))
	fmt.Println(IsBipartite(3, [][2]int{{0, 1}, {1, 2}, {2, 0}}))
}
