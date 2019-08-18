package main

import (
	"fmt"
	"strconv"
	"strings"
)

type LinkedNode struct {
	prev  *LinkedNode
	value int
	next  *LinkedNode
}

var dummy = &LinkedNode{
	prev:  nil,
	value: 0,
	next:  nil,
}

func main() {
	insertX(5)
	insertX(2)
	insertX(3)
	insertX(1)
	deleteX(3) // TODO: deleteがうまくいかない
	//	insertX(6)
	//	deleteX(5)

	display()
}

func display() {
	if dummy.next == nil {
		return
	}

	target := dummy.next
	s := ""
	for {
		s = s + strconv.Itoa(target.value) + " "
		if target.next == nil {
			break
		}
		target = target.next
	}
	fmt.Println(strings.TrimRight(s, " "))
}

func insertX(x int) {
	head := &LinkedNode{
		prev:  dummy,
		value: x,
		next:  dummy.next,
	}

	dummy.next = head
}

func serchX(x int) *LinkedNode {
	var target = dummy.next
	for {
		if target.value == x {
			return target
		}
		if target.next == nil {
			return nil
		}
		target = target.next
	}
}

func deleteX(x int) {
	target := serchX(x)
	if target == nil {
		return
	}

	target.prev.next = target.next
	target.next.prev = target.prev

}

func deleteFirst() {
	next := dummy.next.next
	dummy.next = next
	next.prev = dummy
}

func deleteLast() {
	var target = dummy.next
	for {
		if target.next == nil {
			prev := target.prev
			prev.next = nil
			return
		}
		target = target.next
	}
}
