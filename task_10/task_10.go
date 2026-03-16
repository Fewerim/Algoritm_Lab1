package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	N, K, M int
	g       [25][25]bool
	deg     [25]int

	team   []int
	best   = -1
	answer []int
)

func dfs(start int, inside int, sumDeg int) {
	if len(team) == K {
		cross := sumDeg - 2*inside
		inside2 := M - inside - cross
		score := inside + inside2

		if score > best {
			best = score
			answer = append([]int{}, team...)
		}
		return
	}

	for i := start; i <= N; i++ {
		add := 0
		for _, v := range team {
			if g[v][i] {
				add++
			}
		}
		team = append(team, i)
		dfs(i+1, inside+add, sumDeg+deg[i])
		team = team[:len(team)-1]
	}
}

func main() {
	var N, K, M int
	fmt.Scan(&N, &K, &M)

	for i := 0; i < M; i++ {
		var a, b int
		fmt.Scan(&a, &b)

		g[a][b] = true
		g[b][a] = true

		deg[a]++
		deg[b]++
	}

	dfs(1, 0, 0)

	var sb strings.Builder
	for i, v := range answer {
		if i > 0 {
			sb.WriteByte(' ')
		}
		str := strconv.Itoa(v)
		sb.WriteString(str)
	}
	fmt.Print(sb.String())
}
