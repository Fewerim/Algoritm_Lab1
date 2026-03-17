package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

const (
	MaxSize    int = 1e6
	MaxNum     int = 1e6 + 1
	MinSize        = 3
	MinNum     int = -1e6 - 1
	ErrorValue     = 0
)

func input() []int {
	var size int
	fmt.Scan(&size)
	if size > MaxSize || size < MinSize {
		return []int{}
	}

	inputArr := make([]int, size)

	for i := 0; i < len(inputArr); i++ {
		fmt.Scan(&inputArr[i])
	}

	return inputArr
}

func MaxThree(input []int) int {
	length := len(input)
	if length < 3 {
		return ErrorValue
	}

	max1, max2, max3 := MinNum, MinNum, MinNum
	min1, min2 := MaxNum, MaxNum

	// проходим по всем элементам массива
	for _, num := range input {
		// сдвигаем максимумы
		if num > max1 {
			max3, max2, max1 = max2, max1, num
		} else if num > max2 {
			max3, max2 = max2, num
		} else if num > max3 {
			max3 = num
		}

		// аналогично с минимумами
		if num < min1 {
			min2, min1 = min1, num
		} else if num < min2 {
			min2 = num
		}
	}

	res1 := max1 * max2 * max3
	res2 := min1 * min2 * max1

	return maxNum(res1, res2)
}

func maxNum[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(MaxThree(input()))
}
