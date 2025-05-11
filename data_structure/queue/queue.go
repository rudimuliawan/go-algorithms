package queue

type node struct {
	prev *node
	val  any
}

type Queue struct {
	head   *node
	tail   *node
	Length int
}

func (q *Queue) IsEmpty() bool {
	return q.Length == 0
}

func (q *Queue) Enqueue(val any) {
	newNode := &node{val: val}

	if q.head == nil {
		q.head = newNode
	} else {
		q.tail.prev = newNode
	}
	q.tail = newNode

	q.Length++
}

func (q *Queue) Dequeue() any {
	if q.IsEmpty() {
		return nil
	}

	val := q.head.val

	q.head = q.head.prev
	if q.head == nil {
		q.tail = nil
	}

	q.Length--
	return val
}

func (q *Queue) GetItems() []any {
	var items []any

	curr := q.head
	for curr != nil {
		items = append(items, curr.val)
		curr = curr.prev
	}

	return items
}
