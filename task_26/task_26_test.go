package main

import (
	"fmt"
	"testing"
)

// вспомогательная функция: выполняет операции и возвращает результат X
func runOps(ops []string) []string {
	h := Heap{
		nodes: make([]Node, 0),
		pos:   make([]int, len(ops)+5),
	}

	for i := range h.pos {
		h.pos[i] = -1
	}

	opIndex := 0
	result := []string{}

	for _, op := range ops {
		if op[0] == 'A' {
			var x int
			fmt.Sscanf(op, "A %d", &x)
			h.Push(x, opIndex)

		} else if op[0] == 'X' {
			val, ok := h.Pop()
			if !ok {
				result = append(result, "*")
			} else {
				result = append(result, fmt.Sprintf("%d", val))
			}

		} else if op[0] == 'D' {
			var x, y int
			fmt.Sscanf(op, "D %d %d", &x, &y)
			h.Update(x, y)
		}

		opIndex++
	}

	return result
}

func TestBasic(t *testing.T) {
	ops := []string{
		"A 5",
		"A 3",
		"X",
		"A 10",
		"D 0 1",
		"X",
		"X",
	}

	expected := []string{"3", "1", "10"}
	got := runOps(ops)

	for i := range expected {
		if expected[i] != got[i] {
			t.Errorf("expected %v, got %v", expected, got)
			return
		}
	}
}

func TestEmpty(t *testing.T) {
	ops := []string{
		"X",
		"A 7",
		"D 1 2",
		"X",
		"X",
	}

	expected := []string{"*", "2", "*"}
	got := runOps(ops)

	for i := range expected {
		if expected[i] != got[i] {
			t.Errorf("expected %v, got %v", expected, got)
			return
		}
	}
}

func TestDecreaseKey(t *testing.T) {
	ops := []string{
		"A 10",
		"A 20",
		"A 30",
		"D 1 5",
		"X",
		"X",
	}

	expected := []string{"5", "10"}
	got := runOps(ops)

	for i := range expected {
		if expected[i] != got[i] {
			t.Errorf("expected %v, got %v", expected, got)
			return
		}
	}
}

func TestMultipleUpdates(t *testing.T) {
	ops := []string{
		"A 50",
		"A 40",
		"A 30",
		"D 0 10",
		"D 1 5",
		"X",
		"X",
		"X",
	}

	expected := []string{"5", "10", "30"}
	got := runOps(ops)

	for i := range expected {
		if expected[i] != got[i] {
			t.Errorf("expected %v, got %v", expected, got)
			return
		}
	}
}
