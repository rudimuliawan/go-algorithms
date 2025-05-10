package data_structure

// Node stores information about data
type Node struct {
	val  any
	next *Node
}

type Stack struct {
	top    *Node
	Length int
}

func (s *Stack) Peek() any {
	return s.top.val
}

func (s *Stack) GetItems() []any {
	var data []any
	curr := s.top

	for curr != nil {
		data = append(data, curr.val)
		curr = curr.next
	}

	return data
}

func (s *Stack) Push(value any) {
	newTop := &Node{val: value}
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
