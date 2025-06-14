package stack

type Stack[T any] struct {
	arr []T
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{}
}

func (s *Stack[T]) Push(el T) {
	s.arr = append(s.arr, el)
}

func (s *Stack[T]) Pop() *T {
	if len(s.arr) == 0 {
		return nil
	}
	last := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return &last
}

func (s *Stack[T]) Len() int {
	return len(s.arr)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Len() == 0
}
