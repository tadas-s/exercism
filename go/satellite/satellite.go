package satellite

import (
	"errors"
	"slices"
)

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

func TreeFromTraversals(preorder, inorder []string) (*Node, error) {
	if valid, err := validNodeList(preorder, inorder); !valid {
		return nil, err
	}

	if len(preorder) == 0 {
		return nil, nil
	}

	tree := &Node{preorder[0], nil, nil}

	rootIndex := slices.Index(inorder, preorder[0])

	leftInOrder := inorder[:rootIndex]
	rightInOrder := inorder[(rootIndex + 1):]
	leftPreOrder := preorder[1 : len(leftInOrder)+1]
	rightPreOrder := preorder[len(leftInOrder)+1:]

	tree.Left, _ = TreeFromTraversals(leftPreOrder, leftInOrder)
	tree.Right, _ = TreeFromTraversals(rightPreOrder, rightInOrder)

	return tree, nil
}

func validNodeList(preorder, inorder []string) (bool, error) {
	presence := map[string]bool{}

	if len(preorder) != len(inorder) {
		return false, errors.New("traversals must have the same length")
	}

	for _, node := range preorder {
		if presence[node] {
			return false, errors.New("traversals must contain unique items")
		}
		presence[node] = true
	}

	for _, node := range preorder {
		if slices.Index(inorder, node) == -1 {
			return false, errors.New("traversals must have the same elements")
		}
	}

	return true, nil
}
