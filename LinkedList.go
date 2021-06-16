package main

import (
	"fmt"
	"log"
)

type ListNode struct {
	data interface{}
	next *ListNode
	prev *ListNode
}

type LinkedList struct {
	head *ListNode
	tail *ListNode
	size int
}

func (list *LinkedList) insert(x interface{}) {
	ptr := &ListNode{data : x}
	if list.head == nil {
		list.head, list.tail = ptr, ptr
	} else{
		list.tail.next = ptr
		ptr.prev = list.tail
		list.tail = ptr
	}
	list.size++
}

func (list *LinkedList) get(index int) interface{}{
	place := 0
	curr := list.head
	if curr == nil{
		log.Fatal("ERROR OUT OF BOUNDS")
	}
	if index < 0 {
		return list.head.data
	}
	for curr != nil {
		if place == index{
			return curr.data
		}
		place++
		curr = curr.next
	}
	return list.tail.data
}

func (list *LinkedList) PopEnd() interface{}{
	if list.size == 0 {
		log.Fatal("No more elements")
	}
	list.size--
	if list.head == list.tail {
		temp := list.tail
		list.tail, list.head = nil, nil
		return temp.data
	} else{
		temp := list.tail
		list.tail = list.tail.prev
		list.tail.next = nil
		return temp.data
	}
	return nil
}

func (list *LinkedList) PopFront() interface{}{
	if list.size == 0 {
		log.Fatal("No more elements")
	}
	list.size--
	if list.head == list.tail {
		temp := list.tail
		list.tail, list.head = nil, nil
		return temp.data
	} else{
		temp := list.head
		list.head = list.head.next
		list.head.prev = nil
		return temp.data
	}
	return nil
}

func (list *LinkedList) pushBack(x interface{}) {
	ptr := &ListNode{data : x}
	if list.head == nil {
		list.head, list.tail = ptr, ptr
	} else{
		list.tail.next = ptr
		ptr.prev = list.tail
		list.tail = ptr
	}
	list.size++
}

func (list *LinkedList) pushFront(x interface{}) {
	ptr := &ListNode{data : x}
	if list.head == nil {
		list.head, list.tail = ptr, ptr
	} else{
		list.head.prev = ptr
		ptr.next = list.head
		list.head = ptr
	}
	list.size++
}

func (list *LinkedList) isEmpty() bool{
	if list.size == 0 {
		return true
	} else { 
		return false
	}
}

func main() {
	list := &LinkedList{}
	list.insert(2);
	list.insert(3);
	list.insert(4);
	fmt.Println(list.get(-1));
}