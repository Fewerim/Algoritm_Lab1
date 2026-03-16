package main

// TODO: append / dequeue / уменьшить элемент, добавленный во время одной из операций
//  если очередь пуста вывести *
//  Команды:
//  А х - требуется добавить элемент x в очередь.
//  X - требуется удалить из очереди минимальный элемент и вывести его в выходной файл. (иначе *)
//  D x y - требуется заменить значение элемента, добавленного в очередь операцией A в строке входного файла номер x+1, на y.

type Node struct {
	val int
	idx int
}

type Heap struct {
	nodes []Node
	pos   []int
}

func (h *Heap) Swap(i, j int) {
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]

	h.pos[h.nodes[i].idx] = i
	h.pos[h.nodes[j].idx] = j
}

func (h *Heap) NodeUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2

		if h.nodes[parent].val < h.nodes[i].val {
			break
		}

		h.Swap(i, parent)
		i = parent
	}
}

func (h *Heap) NodeDown(i int) {
	n := len(h.nodes)

	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i

		if left < n && h.nodes[left].val < h.nodes[smallest].val {
			smallest = left
		}

		if right < n && h.nodes[right].val < h.nodes[smallest].val {
			smallest = right
		}

		if smallest != i {
			break
		}

		h.Swap(i, smallest)

		i = smallest
	}
}

func (h *Heap) Push(val, id int) {
	h.nodes = append(h.nodes, Node{val, id})

	i := len(h.nodes) - 1
	h.pos[id] = i

	h.NodeUp(i)
}

func (h *Heap) Pop() (int, bool) {
	if len(h.nodes) == 0 {
		return 0, false
	}

	root := h.nodes[0]

	last := len(h.nodes) - 1
	h.Swap(0, last)

	h.nodes = h.nodes[:last]
	h.pos[root.idx] = -1

	if len(h.nodes) > 1 {
		h.NodeDown(0)
	}

	return root.val, true
}

func (h *Heap) Update(id, newVal int) {
	i := h.pos[id]

	h.nodes[i].val = newVal

	h.NodeUp(i)
}
