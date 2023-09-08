package main

import (
	"fmt"
	"math/rand"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func newList() *linkedList {
	return &linkedList{}
}

func newNode(val int, next *node) *node {
	return &node{data: val,
		next: next}
}

func (l *linkedList) prepend(val int) {
	toNext := l.head
	l.head = newNode(val, toNext)
	l.length++
}

func TraverseList(head *node) {
	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}
}

func ReverseListItt(head *node) *linkedList {
	var pre, cur, next *node = nil, head, nil
	length := 0

	for cur != nil {
		length++
		next = cur.next
		cur.next = pre
		pre = cur
		cur = next
	}
	return &linkedList{head: pre,
		length: length}
}

func ReverseListRec(head *node) *node {
	if head == nil || head.next == nil {
		return head
	}

	r := ReverseListRec(head.next)
	head.next.next = head
	head.next = nil
	return r
}

func (l *linkedList) removeNodeByValue(val int) {
	if l.head.data == val {
		l.head = l.head.next
		l.length--
		return
	}

	prevToDel := l.head.next
	for prevToDel == nil {
		if prevToDel.data != val {
			prevToDel = prevToDel.next
		}
	}
	prevToDel.next = prevToDel.next.next
	l.length--

}

func main() {
	ll := newList()

	for _, v := range rand.Perm(5) {
		ll.prepend(v)
	}
	fmt.Println("----- Prepend -----")
	TraverseList(ll.head)

	ll = ReverseListItt(ll.head)
	fmt.Println("----- Reverse 1 -----")
	TraverseList(ll.head)

	ll.head = ReverseListRec(ll.head)
	fmt.Println("----- Reverse 2 -----")
	TraverseList(ll.head)

	ll.removeNodeByValue(1)
	fmt.Println("----- Remove 1 -----")
	TraverseList(ll.head)

	ll.removeNodeByValue(4)
	fmt.Println("----- Remove 4 -----")
	TraverseList(ll.head)

}
