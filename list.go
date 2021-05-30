package main

type node struct {
	Value int
	Next  *node
}

func search(v int, list *node) *node {
	if list == nil {
		return nil
	}

	if list.Value == v {
		return list
	}
	return search(v, list.Next)
}

func insert(v int, list **node) {
	headPtr := &node{v, *list}
	*list = headPtr
}

func predecessor(v int, list *node) *node {
	if list == nil {
		return nil
	}

	if list.Next != nil && list.Next.Value == v {
		return list
	}
	return predecessor(v, list.Next)
}

func delete(v int, list **node) {
	p := search(v, *list)
	if p != nil {
		pred := predecessor(v, *list)
		if pred != nil {
			pred.Next = p.Next
		} else {
			*list = p.Next
		}
	}
}
