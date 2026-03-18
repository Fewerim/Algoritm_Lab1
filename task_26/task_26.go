package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	val int // значение
	idx int // номер операции А
}

type Heap struct {
	nodes []Node // куча
	pos   []int  // позиция элемента в куче
}

// Swap - меняет местами два элемента в куче и обновляет их позиции в массиве pos
func (h *Heap) Swap(i, j int) {
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]

	h.pos[h.nodes[i].idx] = i
	h.pos[h.nodes[j].idx] = j
}

// NodeUp - просеивание вверх (после добавления или уменьшения значения)
func (h *Heap) NodeUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2

		// если родитель <= -> все ок
		if h.nodes[parent].val < h.nodes[i].val {
			break
		}

		h.Swap(i, parent)
		i = parent
	}
}

// NodeDown - просеивание вниз (после удаления минимума)
func (h *Heap) NodeDown(i int) {
	n := len(h.nodes)

	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i

		// ищем минимального среди родителя и детей
		if left < n && h.nodes[left].val < h.nodes[smallest].val {
			smallest = left
		}

		if right < n && h.nodes[right].val < h.nodes[smallest].val {
			smallest = right
		}

		// выходим, если уже минимум
		if smallest == i {
			break
		}

		h.Swap(i, smallest)
		i = smallest
	}
}

// Push - добавление элемента в кучу
func (h *Heap) Push(val, id int) {
	h.nodes = append(h.nodes, Node{val, id})

	i := len(h.nodes) - 1
	h.pos[id] = i

	h.NodeUp(i)
}

// Pop - извлечение минимума
func (h *Heap) Pop() (int, bool) {
	if len(h.nodes) == 0 {
		return 0, false
	}

	root := h.nodes[0]
	last := len(h.nodes) - 1

	h.Swap(0, last) // переносим последний элемент в корень

	// удаляем последний элемент
	h.nodes = h.nodes[:last]
	h.pos[root.idx] = -1

	// восстановление кучи
	if len(h.nodes) > 1 {
		h.NodeDown(0)
	}

	return root.val, true
}

// Update - уменьшение значения элемента
func (h *Heap) Update(id, newVal int) {
	i := h.pos[id]

	h.nodes[i].val = newVal

	h.NodeUp(i)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	h := Heap{
		nodes: make([]Node, 0),
		pos:   make([]int, n+5),
	}

	for i := range h.pos {
		h.pos[i] = -1
	}
	opIndex := 0

	for i := 0; i < n; i++ {
		var cmd string
		fmt.Fscan(in, &cmd)

		if cmd == "A" {
			var x int
			fmt.Fscan(in, &x)
			h.Push(x, opIndex)
		} else if cmd == "X" {
			val, ok := h.Pop()
			if !ok {
				fmt.Fprintln(out, "*")
			} else {
				fmt.Fprintln(out, val)
			}
		} else if cmd == "D" {
			var x, y int
			fmt.Fscan(in, &x, &y)
			h.Update(x, y)
		}

		opIndex++
	}
}
