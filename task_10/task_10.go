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
		cross := sumDeg - 2*inside    // считаем количество ребер между командами
		inside2 := M - inside - cross // ребра внутри 2 команды
		score := inside + inside2     // сплоченность

		if score > best {
			best = score
			answer = append([]int{}, team...)
		}
		return
	}

	// перебираем кандидатов
	for i := start; i <= N; i++ {
		add := 0

		// сколько связей появляется внутри команды
		for _, v := range team {
			if g[v][i] {
				add++
			}
		}
		team = append(team, i)

		newInside := inside + add    // добавляем количества ребер
		newSumDeg := sumDeg + deg[i] // общее количество ребер выходящих из выбранных вершин
		newStart := i + 1            // сдвигаемся далее, чтобы не натыкаться на дубликаты

		dfs(newStart, newInside, newSumDeg)

		// откатываем
		team = team[:len(team)-1]
	}
}

func main() {
	fmt.Scan(&N, &K, &M)

	for i := 0; i < M; i++ {
		var a, b int
		fmt.Scan(&a, &b)

		// отмечаем знакомство
		g[a][b] = true
		g[b][a] = true

		// увеличиваем степени знакомства
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
