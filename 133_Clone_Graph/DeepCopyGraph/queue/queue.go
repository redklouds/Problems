package queue

import "container/list"

type Queue struct {
	list.List
}

//my linked list queue implementation using container/list
func (q *Queue) EnQueue(data interface{}) {
	q.PushBack(data)
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}
func (q *Queue) DeQueue() interface{} {
	if !q.IsEmpty() {
		val := q.Front()
		q.Remove(val)
		return val.Value
	}
	return nil
}
