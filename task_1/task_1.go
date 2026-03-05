package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAX_NUM = 20_000

// inputReader - чтение входных данных
func inputReader() string {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	return scanner.Text()
}

// printArr - выводит массив
func printArr(arr []int) {
	if len(arr) == 0 {
		fmt.Println()
		return
	}

	s := strings.Builder{}
	for i, v := range arr {
		if i > 0 {
			s.WriteString(" ")
		}
		s.WriteString(strconv.Itoa(v))
	}
	fmt.Println(s.String())
}

// SymmetricDifference - функция, которая находит симметрическую разность
func SymmetricDifference(inputLine string, max int) []int {
	input := strings.TrimSpace(inputLine)
	inputArr := strings.Split(input, " ")

	counters := make([]int, max+1)
	isA := true

	for i := range inputArr {
		num, _ := strconv.Atoi(inputArr[i])

		if num == 0 {
			if isA {
				isA = false
			} else {
				break
			}
			continue
		}

		if isA {
			counters[num]++
		} else {
			counters[num]--
		}
	}

	result := make([]int, 0)
	for num := 1; num <= MAX_NUM; num++ {
		if counters[num] != 0 {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	input := inputReader()
	result := SymmetricDifference(input, MAX_NUM)
	printArr(result)
}
