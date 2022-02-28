package internal

type Queue struct {
	values []interface{}
	size   int
}

func NewQueue() *Queue {
	q := &Queue{}
	q.values = make([]interface{}, 0)
	return q
}

func (q *Queue) IsEmpty() bool {
	return (q.size == 0)
}

func (q *Queue) Enqueue(val interface{}) {
	q.values = append(q.values, val)
	q.size++
}

func (q *Queue) Dequeue() interface{} {
	q.size--
	value := q.values[0]
	q.values = q.values[1:]
	return value
}
