package main

import (
	"fmt"
)

const (
	MAX = 5e4
	LOG = 20
)

var g [MAX][]Edge
var tin, tout, dist [MAX]int
var up [MAX][LOG]int
var timer int

type Edge struct {
	to     int
	weight int
}

func scanner() (int, int) {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n-1; i++ {
		var u, v, w int
		fmt.Scan(&u, &v, &w)

		g[u] = append(g[u], Edge{v, w})
		g[v] = append(g[v], Edge{u, w})
	}

	var m int
	fmt.Scan(&m)
	return n, m
}

func dfs(vertex, parent, len int) {
	tin[vertex] = timer
	timer++
	up[vertex][0] = parent
	dist[vertex] = len

	for i := 1; i < LOG; i++ {
		if up[vertex][i-1] != -1 {
			up[vertex][i] = up[up[vertex][i-1]][i-1]
		} else {
			up[vertex][i] = -1
		}
	}

	for _, edge := range g[vertex] {
		to := edge.to
		w := edge.weight

		if to != parent {
			dfs(to, vertex, len+w)
		}
	}

	tout[vertex] = timer
	timer++
}

func isParent(a, b int) bool {
	return (tin[a] <= tin[b]) && (tin[b] <= tin[a])
}

func LCA(a, b int) int {
	if isParent(a, b) {
		return a
	}

	if isParent(b, a) {
		return b
	}

	for i := LOG - 1; i >= 0; i-- {
		if !isParent(up[a][i], b) {
			a = up[a][i]
		}
	}
	return up[a][0]
}

func main() {
	_, m := scanner()
	dfs(0, 0, 0)

	results := make([]int, m)
	for i := 0; i < m; i++ {
		var v, u int
		fmt.Scan(&v, &u)

		lsa := LCA(v, u)
		res := dist[v] + dist[u] - 2*dist[lsa]
		results[i] = res
	}

	for result := range results {
		fmt.Println(results[result])
	}
}
