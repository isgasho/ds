// queue data structure

package queue

import (
	"github.com/higker/ds"
)

// Queue FIFO first in first out
type Queue struct {
	size        int
	front, rear *ds.DulNode // front 就是最先进去的一个元素
}

func New() *Queue {
	return &Queue{
		size:  0,
		front: nil,
	}
}

func (q *Queue) EnQueue(value interface{}) {
	node := &ds.DulNode{Value: value}
	if q.size == 0 {
		q.front = node
		q.rear = node
		q.size++
		return
	}
	q.rear.Next = node
	node.Perv = q.rear
	q.rear = node
	q.size++
}

func (q *Queue) DeQueue() interface{} {
	front := q.front
	q.front = q.front.Next
	q.size--
	return front.Val()
}
