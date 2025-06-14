package queue

type Queue[T any] struct {
	arr []T
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{}
}

func (q *Queue[T]) Enqueue(el T) {
	q.arr = append(q.arr, el)
}

func (q *Queue[T]) Dequeue() *T {
	if len(q.arr) == 0 {
		return nil
	}
	first := q.arr[0]
	q.arr = q.arr[1:]
	return &first
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.arr) == 0
}
