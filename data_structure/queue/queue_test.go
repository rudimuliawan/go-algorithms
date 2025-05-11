package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := Queue{}

	t.Run("Queue.Enqueue", func(t *testing.T) {
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		actual := q.GetItems()
		expected := []any{1, 2, 3}
		for i := range expected {
			if expected[i] != actual[i] {
				t.Errorf("Invalid Queue.Enqueue operation: got %v when expecting %v", actual[i], expected[i])
			}
		}
	})

	t.Run("Queue.Dequeue", func(t *testing.T) {
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		expected := []any{1, 2, 3}
		expectedLength := []any{2, 1, 0}
		for i := range expected {
			v := q.Dequeue()
			if expected[i] != v {
				t.Errorf("Invalid Queue.Dequeue operation: got %v when expecting %v", expected[i], v)
			}

			l := q.Length
			if expectedLength[i] != l {
				t.Errorf("Invalid Queue.Size() operation: got %v when expecting %v", expectedLength[i], l)
			}
		}
	})
}
