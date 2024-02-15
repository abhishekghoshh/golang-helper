package list

type List[T any] struct {
	datas []*T
}

func New[T any](datas ...*T) List[T] {
	return List[T]{datas: datas}
}

func (list *List[T]) Iterator() Iterator[T] {
	return &ListIterator[T]{
		datas: list.datas,
	}
}
func (list *List[T]) Add(data ...*T) {
	list.datas = append(list.datas, data...)
}

type ListIterator[T any] struct {
	index int
	datas []*T
}

func (li *ListIterator[T]) HasNext() bool {
	return li.index < len(li.datas)
}
func (li *ListIterator[T]) Next() *T {
	if li.HasNext() {
		user := li.datas[li.index]
		li.index++
		return user
	}
	return nil
}
