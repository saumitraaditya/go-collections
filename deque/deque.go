package deque

import "fmt"

type node[T any] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

func newNode[T any](value T) *node[T] {
	return &node[T]{value: value}
}

type deque[T any] struct {
	head   *node[T]
	tail   *node[T]
	Length int
}

func NewDeque[T any]() *deque[T] {
	return &deque[T]{}
}

func (d *deque[T]) AddFront(value T) {
	d.Length++
	if d.head == nil {
		d.head = newNode(value)
		d.tail = d.head
	} else {
		temp := newNode(value)
		temp.next = d.head
		d.head.prev = temp
		d.head = temp
	}
}

func (d *deque[T]) Add(value T) {
	if d.head == nil {
		//Length will be incremented in
		//AddFront
		d.AddFront(value)
	} else {
		d.Length++
		temp := newNode(value)
		temp.prev = d.tail
		d.tail.next = temp
		d.tail = temp
	}
}

func (d *deque[T]) PopFirst() (T, error) {
	if d.Length == 0 {
		var t T
		return t, fmt.Errorf("collection is empty")
	} else {
		val := d.head.value
		if d.head.next != nil {
			d.head.next.prev = nil
		}
		d.head = d.head.next
		d.Length--
		return val, nil
	}
}

func (d *deque[T]) Pop() (T, error) {
	if d.Length == 0 {
		var t T
		return t, fmt.Errorf("collection is empty")
	} else {
		val := d.tail.value
		if d.tail.prev != nil {
			d.tail.prev.next = nil
		}
		d.tail = d.tail.prev
		d.Length--
		return val, nil
	}
}

func (d *deque[T]) ListValues() []T {
	list := make([]T, d.Length)
	temp := d.head
	for i := 0; i < int(d.Length); i++ {
		list[i] = temp.value
		temp = temp.next
	}
	return list
}
