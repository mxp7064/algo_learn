package set

type Set[T comparable] map[T]bool

func NewSet[T comparable](elements ...T) Set[T] {
	s := make(Set[T])
	for _, el := range elements {
		s[el] = true
	}
	return s
}

func (s Set[T]) GetElements() []T {
	var result []T
	for el, _ := range s {
		result = append(result, el)
	}
	return result
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Add(el T) {
	s[el] = true
}

func (s Set[T]) Contains(el T) bool {
	return s[el]
}

func (s Set[T]) Delete(el T) {
	delete(s, el)
}

func (s Set[T]) Copy() Set[T] {
	setCopy := make(Set[T])
	for k := range s {
		setCopy[k] = true
	}
	return setCopy
}

// intersection returns intersections of sets a and b
// we need to return a set which contains elements which are present in both sets
func intersection[T comparable](a, b Set[T]) Set[T] {
	result := make(Set[T])
	for el, _ := range a {
		if b[el] { // if element not present, this expression will return false (zero value for bool)
			result[el] = true
		}
	}
	return result
}

func intersectionMulti[T comparable](sets ...Set[T]) Set[T] {
	if len(sets) == 0 {
		return nil
	}

	result := sets[0]              // first set the result to first set
	for _, set := range sets[1:] { // loop remaining sets and intersect each one
		result = intersection(result, set)
	}

	return result
}

// difference returns difference between sets a and b (a - b)
// we need to return a set which contains all elements which are present in a, but not in b
func difference[T comparable](a, b Set[T]) Set[T] {
	result := make(Set[T])
	for el, _ := range a {
		if !b[el] {
			result[el] = true
		}
	}
	return result
}

// union returns union of a and b sets
// we need to return a set which contains all elements which are present in either a or b
func union[T comparable](a, b Set[T]) Set[T] {
	result := make(Set[T])
	for el, _ := range a {
		result[el] = true
	}
	for el, _ := range b {
		result[el] = true
	}
	return result
}
