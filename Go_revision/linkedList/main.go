package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func(list *LinkedList)display(){
	current:=list.head
	for current!=nil{
		fmt.Printf("%d ->",current.data)
		current=current.next
	}
	fmt.Println("nil")
}

func (list *LinkedList) insert(data int) {
    newNode := &Node{data: data, next: list.head}
    list.head = newNode
}

func main() {

	mylist:=LinkedList{}

	mylist.insert(1)
	mylist.insert(2)
	mylist.insert(3)
	mylist.insert(4)
	mylist.insert(5)

	fmt.Println("Linked List:")
	mylist.display()

}