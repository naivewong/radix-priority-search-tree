package priotree

import (
	// "fmt"
)

type PriorityTreeNode struct {
	parent *PriorityTreeNode
	left   *PriorityTreeNode
	right  *PriorityTreeNode
	start  int
	last   int
}

type PriorityTree struct {
	root *PriorityTreeNode
	bits int
}

func NewPriorityTree(bits int) *PriorityTree {
	return &PriorityTree{
		bits: bits,
	}
}

func (tree *PriorityTree) Insert(left, right int) {
	if tree.root == nil {
		tree.root = &PriorityTreeNode{
			start: left,
			last:  right,
		}
		return
	}

	cur := tree.root
	mask := 1 << (tree.bits - 1)
	for mask != 0 {
		if left == cur.start && right == cur.last {
			return
		}

		if right > cur.last || (right == cur.last && left < cur.start) {
			// Swap insert node with current node.
			tmp := left
			left = cur.start
			cur.start = tmp
			tmp = right
			right = cur.last
			cur.last = tmp
		}

		if mask & left == 0 {
			if cur.left == nil {
				cur.left = &PriorityTreeNode{
					start:  left,
					last:   right,
					parent: cur,
				}
				return
			}
			cur = cur.left
		} else {
			if cur.right == nil {
				cur.right = &PriorityTreeNode{
					start:  left,
					last:   right,
					parent: cur,
				}
				return
			}
			cur = cur.right
		}

		mask >>= 1
	}
}

func (tree *PriorityTree) fillHole(node *PriorityTreeNode) {
	if node.left == nil && node.right == nil {
		if node.parent.left == node {
			node.parent.left = nil
		} else {
			node.parent.right = nil
		}
	} else if node.left == nil && node.right != nil {
		node.start = node.right.start
		node.last = node.right.last
		tree.fillHole(node.right)
	} else if node.left != nil && node.right == nil {
		node.start = node.left.start
		node.last = node.left.last
		tree.fillHole(node.left)
	} else {
		if node.left.last < node.right.last || (node.left.last == node.right.last && node.right.start < node.left.start) {
			node.start = node.right.start
			node.last = node.right.last
			tree.fillHole(node.right)
		} else {
			node.start = node.left.start
			node.last = node.left.last
			tree.fillHole(node.left)
		}
	}
}

func (tree *PriorityTree) Delete(left, right int) bool {
	cur := tree.root
	mask := 1 << (tree.bits - 1)

	for mask != 0 {
		if cur.start == left && cur.last == right {
			tree.fillHole(cur)
			return true
		}

		if mask & left == 0 {
			if cur.left == nil {
				return false
			}
			cur = cur.left
		} else {
			if cur.right == nil {
				return false
			}
			cur = cur.right
		}
		mask >>= 1
	}
	return false
}

func (tree *PriorityTree) FirstOverlap(left, right int) *PriorityTreeNode {
	cur := tree.root

	for {
		if cur.start <= right && cur.last >= left {
			return cur
		}

		if cur.left != nil && cur.left.last >= left {
			cur = cur.left
			continue
		}

		if cur.right != nil && cur.right.last >= left {
			cur = cur.right
			continue
		}

		break
	}
	return nil
}

func (tree *PriorityTree) goLeft(left, right int, cur **PriorityTreeNode) bool {
	if (*cur).left != nil && (*cur).left.last >= left {
		(*cur) = (*cur).left
		return true
	}
	return false
}

func (tree *PriorityTree) goRight(left, right int, cur **PriorityTreeNode) bool {
	if (*cur).right != nil && (*cur).right.last >= left {
		(*cur) = (*cur).right
		return true
	}
	return false
}

func (tree *PriorityTree) NextOverlap(left, right int, cur *PriorityTreeNode) *PriorityTreeNode {
	for {
		for tree.goLeft(left, right, &cur) {
			if cur.start <= right && cur.last >= left {
				return cur
			}
		}

		for !tree.goRight(left, right, &cur) {
			if cur != tree.root && cur.parent.right == cur {
				cur = cur.parent
			}
			if cur == tree.root {
				return nil
			}
			cur = cur.parent
		}

		if cur.start <= right && cur.last >= left {
			return cur
		}
	}
}