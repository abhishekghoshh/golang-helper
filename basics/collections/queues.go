package collections

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() int {
	if len(q.items) < 1 {
		return -1 // returns -1 if the queue is empty
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func Queues() {

}
