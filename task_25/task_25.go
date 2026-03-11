package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	for _, line := range strings.Split(inputLine, " ") {
		num, err := strconv.Atoi(line)
		if err != nil {
			return 0, nil, err
		}
		results = append(results, num)
	}

	return results[0], results[1:], nil
}

func main() {
	n, arr, err := input()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(isHeap(n, arr))
}
