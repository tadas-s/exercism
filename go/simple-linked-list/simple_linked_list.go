package simplelinkedlist

import "slices"

import "errors"

type List struct {
	head *Element
}

type Element struct {
	value int
	next  *Element
}

func New(elements []int) *List {
	list := &List{nil}

	for _, element := range elements {
		list.Push(element)
	}

	return list
}

func (l *List) Size() int {
	count := 0
	head := l.head

	for head != nil {
		count++
		head = head.next
	}

	return count
}

func (l *List) Push(element int) {
	newElement := &Element{element, nil}

	if l.head == nil {
		l.head = newElement
	} else {
		lastElement := l.head

		for lastElement.next != nil {
			lastElement = lastElement.next
		}

		lastElement.next = newElement
	}
}

func (l *List) Pop() (int, error) {
	if l.head == nil {
		return 0, errors.New("list is empty")
	}

	lastElement := l.head
	var newLastElement *Element

	for lastElement.next != nil {
		newLastElement = lastElement
		lastElement = lastElement.next
	}

	if newLastElement != nil {
		newLastElement.next = nil
		return lastElement.value, nil
	}

	l.head = nil
	return lastElement.value, nil
}

func (l *List) Peek() (int, error) {
	if l.head == nil {
		return 0, errors.New("list is empty")
	}

	lastElement := l.head

	for lastElement.next != nil {
		lastElement = lastElement.next
	}

	return lastElement.value, nil
}

func (l *List) Array() []int {
	var values []int

	for element := l.head; element != nil; element = element.next {
		values = append(values, element.value)
	}

	return values
}

func (l *List) Reverse() *List {
	reversed := New([]int{})

	for _, value := range slices.Backward(l.Array()) {
		reversed.Push(value)
	}

	return reversed
}
