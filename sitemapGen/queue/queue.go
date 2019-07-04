package queue

import (
    "container/list"
    "errors"
)

type Queue interface {
    Push(e interface{})
    Pop() (interface{}, error)
    Peek() interface{}
    Size() int
}

type LinkedListQueue struct {
    l *list.List
}

func (q LinkedListQueue) Push(e interface{}) {
    q.l.PushBack(e)
}

func (q LinkedListQueue) Pop() (interface{}, error) {
    f := q.l.Front()
    if f == nil {
        return nil, errors.New("can't pop empty queue")
    }
    q.l.Remove(f)
    return f.Value, nil
}

func (q LinkedListQueue) Peek() interface{} {
    return q.l.Front().Value
}

func (q LinkedListQueue) Size() int {
    return q.l.Len()
}

func NewLinkedListQueue () LinkedListQueue {
    l := list.New()
    return LinkedListQueue{l}
}