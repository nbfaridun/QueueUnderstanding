package main

import "testing"

func TestRemoveFirstAndLastElement(t *testing.T) {
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

	k := queue.length

	queue.removeFirstPerson()

	if k - queue.length != 1 {
		want := k - 1
		got := queue.length
		t.Errorf("method len() in Method( add ) is not coorect for empty queue got %d want %d\n", got, want)
	}

	queue.removeLastPerson()

	if k - queue.length != 2 {
		want := k - 1
		got := queue.length
		t.Errorf("method len() in Method( add ) is not coorect for empty queue got %d want %d\n", got, want)
	}
}