package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	MAX        = 2 * 109
	minN, maxN = 1, 106
	YES        = "YES"
	NO         = "NO"
)

func isHeap(n int, array []int) string {
	if len(array) == 0 {
		return NO
	}

	for i := 0; i < n; i++ {
		left := 2*i + 1
		right := 2*i + 2

		if left < n && array[i] > array[left] {
			return NO
		}
		if right < n && array[i] > array[right] {
			return NO
		}
	}
	return YES
}

func input() (int, []int, error) {
	var n int
	fmt.Scanln(&n)
	if n < minN || n > maxN {
		return 0, nil, errors.New("n out of range")
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	fields := strings.Fields(line)
	results := make([]int, n)
	for i, field := range fields {
		num, _ := strconv.Atoi(field)
		if abs := math.Abs(float64(num)); abs > float64(MAX) {
			return 0, nil, errors.New("num out of range")
		}
		results[i] = num
	}

	if len(results) != n {
		return 0, nil, errors.New("invalid input")
	}

	return n, results, nil
}

func main() {
	n, arr, err := input()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(isHeap(n, arr))
}
