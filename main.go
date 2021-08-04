package main

import (
	"fmt"
)

type List struct {
	begin *ListNode // first Element
	end  *ListNode // last Element
	length  int
}

func (to *List) PrintList() {
	fmt.Printf("QUEUE:\n")

	current := to.begin

	for current != nil {
		fmt.Println(current)
		current = current.Next
	}
	fmt.Printf("\n\n")
}

func (to *List) len() int {
	return to.length
}

func (to *List) Add(node ListNode) {

	if to.len() == 0 {
		to.begin = &node
		to.end = &node
		to.length++
		return
	}

	LastElement := to.end

	to.end.Next = &node
	to.end = &node
	to.end.Prev = LastElement
	to.length++
}

func (to *List) removeFirstPerson() {
	fmt.Printf("\n\nRemoving First Person From Queue...\nDONE!\n\n")

	if to.len() == 0 {
		return
	}

	if to.len() == 1 {
		to.begin = nil
	} else {
		to.begin = to.begin.Next
		to.begin.Prev = nil
	}
	to.length--
}

func (to *List) removeLastPerson() {
	fmt.Printf("\n\nRemoving last person from queue...\nDONE!\n\n")

	if to.len() == 0 {
		return
	}

	if to.len() == 1 {
		to.begin = nil
	} else {
		to.end = to.end.Prev
		to.end.Next = nil
	}
	to.length--
}

type ListNode struct {
	Prev *ListNode
	Name string
	Purchases int
	Next *ListNode
}

func main() {

	queue := List{}
	// Добавление 5 человек
	person := ListNode{
		Prev:      nil,
		Name:      "Surush",
		Purchases: 150,
		Next:      nil,
	}
	queue.Add(person)

	person = ListNode{
		Prev:      nil,
		Name:      "Azam",
		Purchases: 100,
		Next:      nil,
	}
	queue.Add(person)

	person = ListNode{
		Prev:      nil,
		Name:      "Raboni",
		Purchases: 120,
		Next:      nil,
	}
	queue.Add(person)

	person = ListNode{
		Prev:      nil,
		Name:      "Faridun",
		Purchases: 50,
		Next:      nil,
	}
	queue.Add(person)

	person = ListNode{
		Prev:      nil,
		Name:      "Aliakbar",
		Purchases: 200,
		Next:      nil,
	}
	queue.Add(person)

	queue.PrintList()

	queue.removeFirstPerson()
	fmt.Println(queue)
	queue.PrintList()

	queue.removeLastPerson()
	fmt.Println(queue)
	queue.PrintList()

}
