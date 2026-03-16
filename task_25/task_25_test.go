package main

import "testing"

type TestCase struct {
	Name     string
	InputArr []int
	N        int
	Expected string
}

func TestIsHeap(t *testing.T) {
	tests := []TestCase{
		{Name: "Пустая куча", InputArr: []int{}, N: 0, Expected: NO},
		{Name: "Один элемент", InputArr: []int{3}, N: 1, Expected: YES},
		{Name: "Валидная куча", InputArr: []int{3, 4, 7, 8, 13}, N: 5, Expected: YES},
		{Name: "Невалидный корень", InputArr: []int{3, 2, 7, 8, 13}, N: 5, Expected: NO},
		{Name: "Невалидный лист", InputArr: []int{3, 4, 7, 8, 1}, N: 5, Expected: NO},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := isHeap(test.N, test.InputArr)
			if got != test.Expected {
				t.Errorf("isHeap(%v, %v) = %v; expected %v", test.N, test.InputArr, got, test.Expected)
			}
		})
	}
}
