package main

import "testing"

func TestSearch(t *testing.T) {
	next := node{Value: 2}
	first := &node{1, &next}

	got := search(2, first)

	if got.Value != 2 {
		t.Fatalf("search(2, %v) = %v; wanted 2", first, got)
	}
}

func TestInsert(t *testing.T) {
	next := node{Value: 2}
	first := &node{1, &next}

	insert(3, &first)

	got := first.Value

	if got != 3 {
		t.Fatalf("insert(3, %v); wanted inserted value: 3; got: %v", first, got)
	}

	got = first.Next.Value

	if got != 1 {
		t.Fatalf("insert(3, %v); wanted next value: 1; got: %v", first, got)
	}
}

func TestDelete(t *testing.T) {

	t.Run("delete head", func(t *testing.T) {
		next := node{Value: 2}
		first := &node{1, &next}

		delete(1, &first)

		if first.Value != 2 {
			t.Fatalf("delete(1, &%v), wanted head value: 2, got: %v ", *first, first.Value)
		}
	})
	t.Run("delete second element", func(t *testing.T) {
		next := node{Value: 2}
		first := &node{1, &next}

		delete(2, &first)

		if first.Value != 1 {
			t.Fatalf("delte(2, &%v), wanted head value: 1, got: %v", *first, first.Value)
		}

	})

}
