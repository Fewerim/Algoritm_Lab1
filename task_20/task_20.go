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

func readMatrix(in *bufio.Reader, N, M int) [][]int {
	a := make([][]int, N+1)
	for i := range a {
		a[i] = make([]int, M+1) // для каждой строки создаем массив столбцов
	}

	// считываем матрицу
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}
	return a
}

func buildPermutations(a [][]int, N, M int) ([]int, []int) {
	P := make([]int, N+1) // перестановка строк
	Q := make([]int, M+1) // перестановка столбцов

	// определяем строки
	for i := 1; i <= N; i++ {
		P[i] = (a[i][1]-1)/M + 1
	}

	// определяем столбцы
	for j := 1; j <= M; j++ {
		Q[j] = (a[1][j]-1)%M + 1
	}

	return P, Q
}

func fixRows(P []int) []Op {
	ops := []Op{}
	N := len(P) - 1

	// бежим по всем позициям, пока i не на своем месте, добавляем операцию: поменять строки
	for i := 1; i <= N; i++ {
		for P[i] != i {
			j := P[i]
			ops = append(ops, Op{'R', i, j})
			swap(P, i, j)
		}
	}
	return ops
}

func fixCols(Q []int) []Op {
	ops := []Op{}
	M := len(Q) - 1

	// бежим по всем позициям, пока j не на своем месте, добавляем операцию: поменять столбцы
	for j := 1; j <= M; j++ {
		for Q[j] != j {
			k := Q[j]
			ops = append(ops, Op{'C', j, k})
			swap(Q, j, k)
		}
	}
	return ops
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, M int
	fmt.Fscan(in, &N, &M)

	a := readMatrix(in, N, M)
	P, Q := buildPermutations(a, N, M)

	ops := []Op{}
	ops = append(ops, fixRows(P)...)
	ops = append(ops, fixCols(Q)...)

	fmt.Fprintln(out, len(ops))
	for _, op := range ops {
		fmt.Fprintf(out, "%c %d %d\n", op.t, op.x, op.y)
	}
}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}
