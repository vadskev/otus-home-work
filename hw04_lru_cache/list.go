package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	first *ListItem
	last  *ListItem
	len   int
}

func (item *list) Len() int {
	return item.len
}

func (item *list) Front() *ListItem {
	return item.first
}

func (item *list) Back() *ListItem {
	return item.last
}

func (item *list) PushFront(v interface{}) *ListItem {
	item.len++

	tmpItem := new(ListItem)
	tmpItem.Value = v

	tmpItem.Next = item.first
	tmpItem.Prev = nil

	if item.first != nil {
		item.first.Prev = tmpItem
		item.first = tmpItem
	} else {
		item.first = tmpItem
		item.last = tmpItem
	}

	return tmpItem
}

func (item *list) PushBack(v interface{}) *ListItem {
	item.len++

	tmpItem := new(ListItem)
	tmpItem.Value = v

	tmpItem.Next = nil
	tmpItem.Prev = item.last

	if item.last != nil {
		item.last.Next = tmpItem
		item.last = tmpItem
	} else {
		item.last = tmpItem
		item.first = tmpItem
	}

	return tmpItem
}

func (item *list) Remove(i *ListItem) {
	item.len--

	if i == item.first {
		item.first = i.Next
	}
	if i == item.last {
		item.last = i.Prev
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	}

	i.Next = nil
	i.Prev = nil
}

func (item *list) MoveToFront(i *ListItem) {
	switch {
	case i.Prev == nil && i.Next != nil:
		return
	case i.Next == nil && i.Prev != nil:
		item.last.Prev.Next = nil
		item.last = item.last.Prev
	case i.Next == nil && i.Prev == nil:
		return
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}

	item.first.Prev = i

	i.Prev = nil
	i.Next = item.first

	item.first = i
}

func NewList() List {
	return new(list)
}
