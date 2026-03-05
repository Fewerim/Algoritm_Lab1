package main

import (
	"fmt"
	"strings"
	"testing"
)

type TestCase struct {
	Name     string
	Input    string
	Expected []int
}

func New(name, input string, expected []int) *TestCase {
	return &TestCase{Name: name, Input: input, Expected: expected}
}

func createLargeInput(length int) string {
	s := strings.Builder{}

	for i := 0; i < length/2; i++ {
		s.WriteString(fmt.Sprintf("%d ", i))
	}
	s.WriteString("0 ")

	for i := length/2 - 30; i < length-1; i++ {
		s.WriteString(fmt.Sprintf("%d ", i))
	}
	s.WriteString("0 ")

	return s.String()
}

func BenchmarkSymmetricDifference(b *testing.B) {
	length := 19_000
	cases := []*TestCase{
		New("test_1", "1 2 3 4 5 0 1 7 5 8 0", nil),
		New("test_2", "1 2 6 8 7 3 0 4 1 6 2 3 9 0", nil),
		New("large_test_3", createLargeInput(length), nil),
	}

	for _, v := range cases {
		b.Run(v.Name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				SymmetricDifference(v.Input, 20_000)
			}
		})
	}
}
