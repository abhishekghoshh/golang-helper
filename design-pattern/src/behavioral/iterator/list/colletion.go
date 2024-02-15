package list

type Collection[T any] interface {
	Iterator() Iterator[T]
	Add(data ...*T)
}
