package stack

import "testing"

func TestStack(t *testing.T) {
	s := Stack{}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	t.Run("Stack.Push", func(t *testing.T) {
		actual := s.GetItems()
		expected := []any{3, 2, 1}
		for i := range expected {
			if expected[i] != actual[i] {
				t.Errorf("Invalid Stack.Push operation: got %v when expecting %v", actual[i], expected[i])
			}
		}

		if s.Length != len(expected) {
			t.Errorf("Invalid s.Length: got %v when expecting %v", s.Length, len(expected))
		}
	})

	t.Run("Stack.Peek", func(t *testing.T) {
		s := Stack{}
		for i := 1; i <= 3; i++ {
			s.Push(i)
			if s.Peek() != i {
				t.Errorf("Invalid Stack.Peek operation: got %v when expecting %v", s.Peek(), i)
			}
		}
	})

	t.Run("Stack.Pop", func(t *testing.T) {
		expected := []any{3, 2, 1}
		for i := range expected {
			v := s.Pop()
			if v != expected[i] {
				t.Errorf("Invalid Stack.Pop operation : got %v, but expected %v", i, expected[i])
			}
		}
	})
}
