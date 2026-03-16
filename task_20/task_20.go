package main

import (
	"bufio"
	"fmt"
	"os"
)

type Op struct {
	t byte
	x int
	y int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, M int
	fmt.Fscan(in, &N, &M)

	a := make([][]int, N+1)
	for i := range a {
		a[i] = make([]int, M+1)
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}

	P := make([]int, N+1)
	Q := make([]int, M+1)

	for i := 1; i <= N; i++ {
		P[i] = (a[i][1]-1)/M + 1
	}

	for j := 1; j <= M; j++ {
		Q[j] = (a[1][j]-1)%M + 1
	}

	ops := make([]Op, 0)

	for i := 1; i <= N; i++ {
		for P[i] != i {
			j := P[i]
			ops = append(ops, Op{'R', i, j})
			P[i], P[j] = P[j], P[i]
		}
	}

	for j := 1; j <= M; j++ {
		for Q[j] != j {
			k := Q[j]
			ops = append(ops, Op{'C', j, k})
			Q[j], Q[k] = Q[k], Q[j]
		}
	}

	fmt.Fprintln(out, len(ops))
	for _, op := range ops {
		fmt.Fprintf(out, "%c %d %d\n", op.t, op.x, op.y)
	}
}
