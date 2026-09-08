package pov

import "slices"

type Tree struct {
	value    string
	parent   *Tree
	children []*Tree
}

// New creates and returns a new Tree with the given root value and children.
func New(value string, children ...*Tree) *Tree {
	tree := &Tree{
		value:    value,
		parent:   nil,
		children: children,
	}

	for _, child := range tree.children {
		child.parent = tree
	}

	return tree
}

// Value returns the value at the root of a tree.
func (tr *Tree) Value() string {
	return tr.value
}

// Children returns a slice containing the children of a tree.
// There is no need to sort the elements in the result slice,
// they can be in any order.
func (tr *Tree) Children() []*Tree {
	return tr.children
}

// String describes a tree in a compact S-expression format.
// This helps to make test outputs more readable.
// Feel free to adapt this method as you see fit.
func (tr *Tree) String() string {
	if tr == nil {
		return "nil"
	}
	result := tr.Value()
	if len(tr.Children()) == 0 {
		return result
	}
	for _, ch := range tr.Children() {
		result += " " + ch.String()
	}
	return "(" + result + ")"
}

// Subtree finds node with `value` if exists
func (tr *Tree) Subtree(value string) *Tree {
	if tr.value == value {
		return tr
	}

	for _, child := range tr.children {
		subtree := child.Subtree(value)

		if subtree != nil {
			return subtree
		}
	}

	return nil
}

// POV problem-specific functions

// FromPov returns the pov from the node specified in the argument.
func (tr *Tree) FromPov(from string) *Tree {
	subtree := tr.Subtree(from)

	if subtree == nil {
		return nil
	}

	pov := subtree.Clone()

	if subtree.parent != nil {
		pov.children = append(pov.children, subtree.parent.CloneUp(subtree))
	}

	return pov
}

// Clone returns a copy of a given tree (or subtree)
func (tr *Tree) Clone() *Tree {
	var children []*Tree

	for _, child := range tr.children {
		children = append(children, child.Clone())
	}

	tree := New(tr.value, children...)

	return tree
}

// CloneUp returns an "upside down" copy of a subtree
// excluding the `except` node which becomes the new root
func (tr *Tree) CloneUp(except *Tree) *Tree {
	var children []*Tree

	for _, child := range tr.children {
		if child == except {
			continue
		}

		children = append(children, child.Clone())
	}

	if tr.parent != nil {
		children = append(children, tr.parent.CloneUp(tr))
	}

	return New(tr.value, children...)
}

// PathTo returns the shortest path between two nodes in the tree.
func (tr *Tree) PathTo(from, to string, visited ...string) []string {
	var subtree *Tree

	if tr.value == from {
		subtree = tr
	} else {
		subtree = tr.Subtree(from)
	}

	if subtree == nil {
		return nil
	}

	visited = append(visited, subtree.value)

	if from == to {
		return visited
	}

	if subtree.parent != nil && !slices.Contains(visited, subtree.parent.value) {
		viaParent := subtree.parent.PathTo(subtree.parent.value, to, visited...)

		if viaParent != nil {
			return viaParent
		}
	}

	for _, child := range subtree.children {
		if !slices.Contains(visited, child.value) {
			viaChild := child.PathTo(child.value, to, visited...)

			if viaChild != nil {
				return viaChild
			}
		}
	}

	return nil
}
