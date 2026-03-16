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
)

/*
TODO:
  - неубывающая пирамида на основе массива
    Условия для каждого 1<=i<=n:
  - if 2i <= n then a[i] <= a[2i]
  - if 2i+1 <= n then a[i] <= a[2i+1]
*/
func isHeap(n int, array []int) bool {
	if len(array) == 0 {
		return false
	}

	for i := 0; i < n; i++ {
		left := 2 * i
		right := 2*i + 1

		if left < n && array[i] > array[left] {
			return false
		}
		if right < n && array[i] > array[right] {
			return false
		}
	}
	return true
}

func input() (int, []int, error) {
	results := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputLine := scanner.Text()
	var n int
	for i, line := range strings.Split(inputLine, " ") {
		if i == 0 {
			n, _ = strconv.Atoi(line)
			if n < minN || n > maxN {
				return 0, nil, errors.New("n out of range")
			}
		}
		num, _ := strconv.Atoi(line)
		if abs := math.Abs(float64(num)); abs > MAX {
			return 0, nil, errors.New("input num invalid")
		}
		results = append(results, num)
	}

	return n, results[1:], nil
}

func main() {
	n, arr, err := input()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(isHeap(n, arr))
}
