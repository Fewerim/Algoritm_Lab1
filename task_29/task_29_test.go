package main

import "testing"

func buildTree(n int, edges [][3]int) *Tree {
	t := NewTree(n)
	for _, e := range edges {
		t.addEdge(e[0], e[1], e[2])
	}
	t.dfs(0, -1, 0)
	return t
}

func TestSingleNode(t *testing.T) {
	tree := buildTree(1, nil)

	if tree.distance(0, 0) != 0 {
		t.Errorf("expected 0, got %d", tree.distance(0, 0))
	}
}

func TestTwoNodes(t *testing.T) {
	tree := buildTree(2, [][3]int{
		{0, 1, 7},
	})

	if tree.distance(0, 1) != 7 {
		t.Errorf("expected 7, got %d", tree.distance(0, 1))
	}
}

func TestStar(t *testing.T) {
	tree := buildTree(5, [][3]int{
		{0, 1, 1},
		{0, 2, 2},
		{0, 3, 3},
		{0, 4, 4},
	})

	if tree.distance(1, 2) != 3 {
		t.Errorf("expected 3, got %d", tree.distance(1, 2))
	}
}

func TestBalanced(t *testing.T) {
	tree := buildTree(7, [][3]int{
		{0, 1, 1},
		{0, 2, 1},
		{1, 3, 2},
		{1, 4, 2},
		{2, 5, 3},
		{2, 6, 3},
	})

	if tree.distance(3, 5) != 7 {
		t.Errorf("expected 7, got %d", tree.distance(3, 5))
	}
}

func TestSameNode(t *testing.T) {
	tree := buildTree(3, [][3]int{
		{0, 1, 5},
		{1, 2, 6},
	})

	if tree.distance(2, 2) != 0 {
		t.Errorf("expected 0, got %d", tree.distance(2, 2))
	}
}

func TestZeroWeights(t *testing.T) {
	tree := buildTree(4, [][3]int{
		{0, 1, 0},
		{1, 2, 0},
		{2, 3, 0},
	})

	if tree.distance(0, 3) != 0 {
		t.Errorf("expected 0, got %d", tree.distance(0, 3))
	}
}
