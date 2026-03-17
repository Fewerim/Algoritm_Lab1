package main

import (
	"bufio"
	"fmt"
	"os"
)

const LOG = 20

type Edge struct {
	to int
	w  int
}

type Tree struct {
	n int

	g [][]Edge

	tin, tout []int
	dist      []int

	up [][]int

	timer int
}

func NewTree(n int) *Tree {
	t := &Tree{
		n:    n,
		g:    make([][]Edge, n),
		tin:  make([]int, n),
		tout: make([]int, n),
		dist: make([]int, n),
		up:   make([][]int, n),
	}

	for i := 0; i < n; i++ {
		t.up[i] = make([]int, LOG)
		for j := 0; j < LOG; j++ {
			t.up[i][j] = -1
		}
	}

	return t
}

// addEdge - добавление ребра
func (t *Tree) addEdge(u, v, w int) {
	t.g[u] = append(t.g[u], Edge{v, w})
	t.g[v] = append(t.g[v], Edge{u, w})
}

// dfs - для подготовки LCA
func (t *Tree) dfs(v, parent, distFromRoot int) {
	t.tin[v] = t.timer
	t.timer++

	t.up[v][0] = parent
	t.dist[v] = distFromRoot

	// бинарный подъём
	for i := 1; i < LOG; i++ {
		if t.up[v][i-1] != -1 {
			t.up[v][i] = t.up[t.up[v][i-1]][i-1]
		}
	}

	for _, e := range t.g[v] {
		if e.to != parent {
			t.dfs(e.to, v, distFromRoot+e.w)
		}
	}

	t.tout[v] = t.timer
	t.timer++
}

// isAncestor - является ли a предком b
func (t *Tree) isAncestor(a, b int) bool {
	return t.tin[a] <= t.tin[b] && t.tout[b] <= t.tout[a]
}

// lca - наименьший общий предок
func (t *Tree) lca(a, b int) int {
	if t.isAncestor(a, b) {
		return a
	}
	if t.isAncestor(b, a) {
		return b
	}

	for i := LOG - 1; i >= 0; i-- {
		if t.up[a][i] != -1 && !t.isAncestor(t.up[a][i], b) {
			a = t.up[a][i]
		}
	}

	return t.up[a][0]
}

// distance - расстояние между вершинами
func (t *Tree) distance(a, b int) int {
	lca := t.lca(a, b)
	return t.dist[a] + t.dist[b] - 2*t.dist[lca]
}

func main() {
	var n int
	fmt.Scan(&n)

	tree := NewTree(n)

	for i := 0; i < n-1; i++ {
		var u, v, w int
		fmt.Scan(&u, &v, &w)
		tree.addEdge(u, v, w)
	}

	// запускаем DFS от корня
	root := 0
	tree.dfs(root, -1, 0)

	var m int
	fmt.Scan(&m)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// обрабатываем запросы
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)

		res := tree.distance(u, v)
		fmt.Fprintln(writer, res)
	}
}
