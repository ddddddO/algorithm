package main

import (
	"bufio"
	"fmt"
	"os"
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
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}

	for i := 0; i < n; i++ {
		sc.Scan()
		text := sc.Text()

		switch text {
		case "deleteFirst":
			deleteFirst()

		case "deleteLast":
			deleteLast()

		default:
			splited := []string{}

			if strings.Contains(text, "delete") {
				splited = strings.Split(text, " ")
				n, err := strconv.Atoi(splited[1])
				if err != nil {
					panic(err)
				}
				deleteX(n)
			}
			if strings.Contains(text, "insert") {
				splited = strings.Split(text, " ")
				n, err := strconv.Atoi(splited[1])
				if err != nil {
					panic(err)
				}
				insertX(n)
			}
		}
	}

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

	// 以下、if文腑に落ちない
	if dummy.next != nil {
		dummy.next.prev = head
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

	if target.next != nil {
		target.next.prev = target.prev
	}

}

func deleteFirst() {
	next := dummy.next.next
	dummy.next = next

	if next != nil {
		next.prev = dummy
	}
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
