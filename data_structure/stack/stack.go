package stack

// node stores information about data
type node struct {
	val  any
	next *node
}

type Stack struct {
	top    *node
	Length int
}

func (s *Stack) Peek() any {
	return s.top.val
}

func (s *Stack) GetItems() []any {
	var items []any
	curr := s.top

	for curr != nil {
		items = append(items, curr.val)
		curr = curr.next
	}

	return items
}

func (s *Stack) Push(value any) {
	newTop := &node{val: value}
	if s.top == nil {
		s.top = newTop
	} else {
		newTop.next = s.top
		s.top = newTop
	}

	s.Length++
}

func (s *Stack) Pop() any {
	val := s.top.val
	if s.top.next != nil {
		s.top = s.top.next
	}

	s.Length--
	return val
}

func (s *Stack) IsEmpty() bool {
	return s.Length == 0
}
