package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	N, K, M int
	g       [25][25]bool
	deg     [25]int

	team   []int
	best   = -1
	answer []int
)

func dfs(start int, inside1 int, sumDeg int) {

	if len(team) == K {

		cross := sumDeg - 2*inside1
		inside2 := M - inside1 - cross
		score := inside1 + inside2

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

		dfs(i+1, inside1+add, sumDeg+deg[i])

		team = team[:len(team)-1]
	}
}

func main() {

	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &N, &K, &M)

	for i := 0; i < M; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)

		g[a][b] = true
		g[b][a] = true

		deg[a]++
		deg[b]++
	}

	dfs(1, 0, 0)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for i, v := range answer {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, v)
	}
}
