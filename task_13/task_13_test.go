package main

import "testing"

type TestCase struct {
	Name     string
	Input    []int
	Expected int
}

func New(name string, input []int, expected int) TestCase {
	return TestCase{
		Name:     name,
		Input:    input,
		Expected: expected,
	}
}

func largeArray(size int) []int {
	arr := make([]int, size)

	for i := 0; i < size; i++ {
		arr[i] = i
	}
	return arr
}

func BenchmarkMaxThree(b *testing.B) {
	cases := []TestCase{
		New("Базовый тест", []int{-1, 2, 3, -4, -2, 5, -1, 5, -3, -2}, 75),
		New("Пустой массив", []int{}, 0),
		New("Большие числа", []int{-1000000, 999999, -999999, 1000000, 0, -500000, 500000}, 999999000000000),
		New("Большой массив", largeArray(1000000), 999994000010999994),
	}

	for _, testCase := range cases {
		b.Run(testCase.Name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				MaxThree(testCase.Input)
			}
		})
	}
}
