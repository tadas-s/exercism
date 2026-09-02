package linkedlist

import "errors"

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.
type Node struct {
	Value any
	next  *Node
	prev  *Node
}

type List struct {
	first *Node
	last  *Node
}

func NewList(elements ...any) *List {
	list := List{
		first: nil,
		last:  nil,
	}

	for _, el := range elements {
		list.Push(el)
	}

	return &list
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v any) {
	node := &Node{
		Value: v,
		next:  nil,
		prev:  nil,
	}

	if l.first == nil {
		l.first = node
		l.last = node
	} else {
		node.next = l.first
		node.next.prev = node
		l.first = node
	}
}

func (l *List) Push(v any) {
	node := &Node{
		Value: v,
		next:  nil,
		prev:  nil,
	}

	if l.last == nil {
		l.last = node
		l.first = node
	} else {
		node.prev = l.last
		node.prev.next = node
		l.last = node
	}
}

func (l *List) Shift() (any, error) {
	if l.first == nil {
		return nil, errors.New("empty list")
	}

	node := l.first
	l.first = node.next

	if l.first != nil {
		l.first.prev = nil
	} else {
		l.last = nil
	}

	return node.Value, nil
}

func (l *List) Pop() (any, error) {
	if l.last == nil {
		return nil, errors.New("empty list")
	}

	node := l.last
	l.last = node.prev

	if l.last != nil {
		l.last.next = nil
	} else {
		l.first = nil
	}

	return node.Value, nil
}

func (l *List) Reverse() {
	for node := l.first; node != nil; node = node.prev {
		node.prev, node.next = node.next, node.prev
	}

	l.first, l.last = l.last, l.first
}

func (l *List) First() *Node {
	return l.first
}

func (l *List) Last() *Node {
	return l.last
}

func (l *List) Count() int {
	count := 0

	for node := l.first; node != nil; node = node.next {
		count++
	}

	return count
}

// Delete removes the first node in a list with a given value.
// Returns true if a node was removed.
func (l *List) Delete(v any) bool {
	found := (*Node)(nil)

	for node := l.first; node != nil && found == nil; node = node.next {
		if node.Value == v {
			found = node
		}
	}

	if found == nil {
		return false
	}

	if l.first == found {
		_, _ = l.Shift()
	} else if l.last == found {
		_, _ = l.Pop()
	} else {
		found.prev.next = found.next
		found.next.prev = found.prev
	}

	return true
}
