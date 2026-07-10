package main

type Edge struct {
	to   int
	sign int
}

func bfs(start int, n int, adj [][]Edge) int {
	count := 0
	visited := make([]bool, n)

	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, edge := range adj[node] {
			if !visited[edge.to] {
				count += edge.sign
				visited[edge.to] = true
				queue = append(queue, edge.to)
			}
		}
	}

	return count
}

func minReorder(n int, connections [][]int) int {
	adj := make([][]Edge, n)

	for _, connection := range connections {
		u := connection[0]
		v := connection[1]

		adj[u] = append(adj[u], Edge{to: v, sign: 1})
		adj[v] = append(adj[v], Edge{to: u, sign: 0})
	}

	return bfs(0, n, adj)
}

/*
type Edge struct {
	to   int
	sign int
}

func dfs(node, parent int, adj [][]Edge, count *int) {
	for _, edge := range adj[node] {
		if edge.to != parent {
			*count += edge.sign
			dfs(edge.to, node, adj, count)
		}
	}
}

func minReorder(n int, connections [][]int) int {
	adj := make([][]Edge, n)

	for _, connection := range connections {
		u := connection[0]
		v := connection[1]

		adj[u] = append(adj[u], Edge{to: v, sign: 1})
		adj[v] = append(adj[v], Edge{to: u, sign: 0})
	}

	count := 0
	dfs(0, -1, adj, &count)

	return count
}
*/


func main(){

}