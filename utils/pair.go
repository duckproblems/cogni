package utils

type Pair[T any] struct {
	Left, Right T
}

func (p Pair[T]) Tuple() (T, T) {
	return p.Left, p.Right
}
