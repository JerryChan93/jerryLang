package commom

import (
	"fmt"
)

type List[T comparable] struct {
	head, tail *node[T]
	len        int
}

type node[T comparable] struct {
	data T
	prev *node[T]
	next *node[T]
}

func (self *List[T]) IsEmpty() bool {
	return self.head == nil
}

func (self *List[T]) Push(data T) {
	n := &node[T]{
		data: data,
		prev: self.tail,
		next: nil,
	}
	if self.head == nil {
		self.head = n
		self.tail = n
		self.len = 0
	}
	self.tail.next = n
	self.tail = n
	self.len += 1
}

func (self *List[T]) Del(data T) {
	for p := self.head; p != nil; p = p.next {
		if data != p.data {
			continue
		}
		if p == self.head {
			self.head = p.next
		}
		if p == self.tail {
			self.tail = p.prev
		}
		if p.prev != nil {
			p.prev.next = p.next
		}
		if p.next != nil {
			p.next.prev = p.prev
		}
		return
	}
}

func (l *List[T]) Each(f func(T)) {
	if l.IsEmpty() {
		fmt.Println("the link list is empty.")
		return
	}
	for p := l.head; p != nil; p = p.next {
		f(p.data)
	}
}

func (l *List[T]) Log() {
	if l.IsEmpty() {
		fmt.Println("the link list is empty.")
		return
	}
	for p := l.head; p != nil; p = p.next {
		fmt.Printf("[%v] -> ", p.data)
	}
	fmt.Println("nil")
}
